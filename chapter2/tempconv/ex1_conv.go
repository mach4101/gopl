package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToK(c Celsius) Kelvim {
	return Kelvim(c - AbsoluteZeroC)
}

func FToK(f Fahrenheit) Kelvim {
	return Kelvim(CToK(FToC(f)))
}

func KToC(k Kelvim) Celsius {
	return Celsius(k + Kelvim(AbsoluteZeroC))
}

func KToF(k Kelvim) Fahrenheit {
	return CToF(KToC(k))
}
