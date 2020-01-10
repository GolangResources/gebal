package gebales

import (
	"fmt"
	"strconv"
)
// Based on codes here: https://es.wikipedia.org/wiki/C%C3%B3digo_cuenta_cliente

// CCC control digit generator 
func cdFactor(digit string) (result string, err error) {
	if (len (digit) != 10) {
		return "00", fmt.Errorf("Incorrect lenght")
	}
	factors := []int{1, 2, 4, 8, 5, 10, 9, 7, 3, 6}

	var factorndigitsum int

	for factorId, factorInt := range factors {
		digitTmpInt , _ := strconv.Atoi(digit[factorId:factorId+1])
		tmpd :=  digitTmpInt * factorInt
		factorndigitsum = factorndigitsum + tmpd
	}
	resint := 11 - factorndigitsum % 11
	if (resint == 10) {
		resint = 1
	} else if (resint == 11) {
		resint = 0
	}
	return strconv.Itoa(resint), nil
}

func CCCCdGen(bank string, branch string, account string) (result string, err error) {
	if (len(bank) != 4) {
		return "", fmt.Errorf("Incorrect bank lenght")
	}
	if (len(branch) != 4) {
		return "", fmt.Errorf("Incorrect branch lenght")
	}
	if (len(account) != 10) {
		return "", fmt.Errorf("Incorrect account lenght")
	}
	cdf1R, cdf1E := cdFactor("00" + bank + branch)
	cdf2R, cdf2E := cdFactor(account)
	if (cdf1E == nil && cdf2E == nil) {
		return cdf1R + cdf2R, nil
	} else {
		return "", fmt.Errorf(cdf1E.Error(), cdf2E.Error())
	}
}

// CCC (Codigo Cuenta Cliente) Checker
func CCCCdCheck(bank string, branch string, controlDigit string, account string) (result bool, err error) {
	if (len(bank) != 4) {
		return false, fmt.Errorf("Incorrect bank lenght")
	}
	if (len(branch) != 4) {
		return false, fmt.Errorf("Incorrect branch lenght")
	}
	if (len(controlDigit) != 2) {
		return false, fmt.Errorf("Incorrect controlDigit lenght")
	}
	if (len(account) != 10) {
		return false, fmt.Errorf("Incorrect account lenght")
	}
	generatedControlDigit, generatedControlDigitErr := CCCCdGen(bank, branch, account)
	if (generatedControlDigitErr == nil) {
		if (generatedControlDigit == controlDigit) {
			return true, nil
		} else {
			return false, nil
		}
	} else {
		return false, generatedControlDigitErr
	}
}
