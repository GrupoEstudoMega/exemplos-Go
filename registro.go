package main

import (
	"fmt"
	"github.com/luisiturrios/gowin"
	"log"
)

func main() {
	chave := `BDS.PasFile\shell\Calcular métricas\Command`
	v, err := gowin.GetReg("HKCR", chave, "")
	//err := gowin.WriteStringReg("HKCR", `BDS.PasFile\shell\Metric\Command`, "", `c:\projetos\go\auditoria-e-metricas\client\clientAudit.exe -arquivo=%1 -browser`)
	if err != nil {
		gowin.WriteStringReg("HKCR", chave, "", `c:\projetos\go\auditoria-e-metricas\client\clientAudit.exe -arquivo=%1 -browser`)
	}
}
