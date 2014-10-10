package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	path := `C:\Mega\megaempresarial\Fontes\FinCPA\uFormAprovaCpagar.pas`
	stat, _ := os.Stat(path)

	fmt.Println(stat.Name())
	short, _ := toShort(path)
	fmt.Println(toLong(short))
}

func toLong(path string) (string, error) {
	p, err := syscall.UTF16FromString(path)
	if err != nil {
		return "", err
	}
	b := p // GetLongPathName says we can reuse buffer
	n, err := syscall.GetLongPathName(&p[0], &b[0], uint32(len(b)))
	if err != nil {
		return "", err
	}
	if n > uint32(len(b)) {
		b = make([]uint16, n)
		n, err = syscall.GetLongPathName(&p[0], &b[0], uint32(len(b)))
		if err != nil {
			return "", err
		}
	}
	b = b[:n]
	return syscall.UTF16ToString(b), nil
}

func toShort(path string) (string, error) {
	p, err := syscall.UTF16FromString(path)
	if err != nil {
		return "", err
	}
	b := p // GetShortPathName says we can reuse buffer
	n, err := syscall.GetShortPathName(&p[0], &b[0], uint32(len(b)))
	if err != nil {
		return "", err
	}
	if n > uint32(len(b)) {
		b = make([]uint16, n)
		n, err = syscall.GetShortPathName(&p[0], &b[0], uint32(len(b)))
		if err != nil {
			return "", err
		}
	}
	return syscall.UTF16ToString(b), nil
}
