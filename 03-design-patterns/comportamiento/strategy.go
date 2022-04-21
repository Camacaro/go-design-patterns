package main

import "fmt"

// Clases Principales (interfaces)
type PasswordProtected struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm
}

// Esta es la firma para la comunicacion entre mi clase principal y las clases que la implementan
type HashAlgorithm interface {
	Hash(p *PasswordProtected)
}

func NewPasswordProtected(user string, passwordName string, hash HashAlgorithm) *PasswordProtected {
	return &PasswordProtected{
		user:          user,
		passwordName:  passwordName,
		hashAlgorithm: hash,
	}
}

func (p *PasswordProtected) SetHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgorithm = hash
}

func (p *PasswordProtected) Hash() {
	p.hashAlgorithm.Hash(p)
}

// END

// Primer type
type SHA struct{}

func (s *SHA) Hash(p *PasswordProtected) {
	fmt.Printf("Hashing %s with SHA\n", p.passwordName)
}

// END

// Segundo type
type MD5 struct{}

func (m *MD5) Hash(p *PasswordProtected) {
	fmt.Printf("Hashing %s with MD5\n", p.passwordName)
}

/*
	Con este patron lo que podemos hacer es cambiar el comportamiento
	interno que tiene una clase (struct) sin modificar su codigo fuente
	ya que se puede crear otras clases (struct sha, md5) e implementar
	internamente en la clase principal (struct PasswordProtected)
*/
func main() {
	sha := &SHA{}
	md5 := &MD5{}

	p1 := NewPasswordProtected("user", "gmail password", sha)
	p1.Hash()
	p1.SetHashAlgorithm(md5)
	p1.Hash()
}
