package vernamotp_cipher

import "fmt"

type Pad struct {
	pages       [][]byte
	currentPage int
}

// NewPad создает новый "одноразовый блокнот", принимая произвольные байты из
// пользователь и настройка страниц от него. Этот метод также позволяет передать
// значение указателя страницы, чтобы упростить возобновление работы с существующей панелью.
func NewPad(material []byte, pageSize int, startPage int) (*Pad, error) {
	// A zero-length page would cause this routine to loop infinitely
	if pageSize < 1 {
		return nil, fmt.Errorf("otp: page length must be greater than 0")
	}

	if len(material) < pageSize {
		return nil, fmt.Errorf("otp: page size too large for pad material")
	}

	// Do the page-splitting work up front
	var pages [][]byte
	for i := 0; i+pageSize <= len(material); i += pageSize {
		pages = append(pages, material[i:i+pageSize])
	}

	// Create the new OTP pad
	p := Pad{
		pages:       pages,
		currentPage: startPage,
	}

	// Set the page index in the new pad
	if err := p.SetPage(startPage); err != nil {
		return nil, err
	}

	return &p, nil
}

// TotalPages возвращает количество страниц в блокноте
func (p *Pad) TotalPages() int {
	return len(p.pages)
}

// RemainingPages возвращает количество неиспользованных страниц в блокноте
func (p *Pad) RemainingPages() int {
	return len(p.pages) - p.currentPage
}

// CurrentPage возвращает текущую позицию указателя страницы
func (p *Pad) CurrentPage() int {
	return p.currentPage
}

// getPage возвращает полезную нагрузку текущей страницы
func (p *Pad) getPage() []byte {
	return p.pages[p.currentPage-1]
}

// SetPage установит указатель страницы
func (p *Pad) SetPage(page int) error {
	if page < 1 || page > p.TotalPages() {
		return fmt.Errorf("otp: page %d out of bounds", page)
	}
	p.currentPage = page
	return nil
}

// NextPage переместит указатель страницы
func (p *Pad) NextPage() error {
	if p.RemainingPages() == 0 {
		return fmt.Errorf("otp: pad exhausted")
	}
	p.currentPage++
	return nil
}

// Шифрование возьмет байтовый срез и будет использовать модульное сложение для шифрования
// полезная нагрузка с использованием текущей страницы.
func (p *Pad) Encrypt(payload []byte) ([]byte, error) {
	page := p.getPage()

	// Page must be at least as long as plain text
	if len(page) < len(payload) {
		return nil, fmt.Errorf("otp: insufficient page size")
	}

	result := make([]byte, len(payload))

	for i := 0; i < len(payload); i++ {
		plainText := int(payload[i])
		secretKey := int(page[i])
		cipherText := (plainText + secretKey) % 255
		result[i] = byte(cipherText)
	}

	return result, nil
}

// Decrypt примет фрагмент байта и обратит процесс Encode к
// перевести зашифрованный текст обратно в необработанные байты. Требуется, чтобы страница
// указатель должен быть установлен в ту же позицию, что и во время Encode().
func (p *Pad) Decrypt(payload []byte) ([]byte, error) {
	page := p.getPage()

	// Page must be at least as long as plain text
	if len(page) < len(payload) {
		return nil, fmt.Errorf("otp: insufficient page size")
	}

	result := make([]byte, len(payload))

	for i := 0; i < len(payload); i++ {
		cipherText := int(payload[i])
		secretKey := int(page[i])
		plainText := (cipherText - secretKey) % 255
		if plainText < 0 {
			plainText += 255
		}
		result[i] = byte(plainText)
	}

	return result, nil
}
