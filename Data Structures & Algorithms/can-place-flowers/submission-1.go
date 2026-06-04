func canPlaceFlowers(flowerbed []int, n int) bool {
    flowerLeft := n

	for i:=0; i<len(flowerbed) && flowerLeft > 0; i++ {
		leftEmpty := i == 0 || i > 0 && flowerbed[i-1] == 0
		rightEmpty := i == len(flowerbed) - 1 || i < len(flowerbed)-1 && flowerbed[i+1] == 0
	
		if flowerbed[i] == 0 && leftEmpty && rightEmpty {
			flowerbed[i] = 1
			flowerLeft--
		}
	}

	return flowerLeft == 0
}