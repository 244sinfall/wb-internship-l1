package main

func task19(toReverse string) string {
	inRunes := []rune(toReverse)
	for idx, lastIdx := 0, len(inRunes)-1; idx < lastIdx; idx, lastIdx = idx+1, lastIdx-1 {
		inRunes[idx], inRunes[lastIdx] = inRunes[lastIdx], inRunes[idx]
	}
	return string(inRunes)
}
