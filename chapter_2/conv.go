package chapter_2

func CToF(c Cels) Fahr { return Fahr(c*9/5 + 32) }
func FToC(f Fahr) Cels { return Cels((f - 32) * 5 / 9) }
