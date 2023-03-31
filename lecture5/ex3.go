func totalPopulation(cities map[string]int) int {
	var cmt int
	for _,v := range cities{
		cmt+=v
	}
	return cmt
}
func mostPopulatedCity(cities map[string]int) string {
	var max int =-99999999
	for _,v := range cities{
		if v > max {
			max = v
		}
	}
	for k,v := range cities{
		if max==v{
			return k
		}
	}
	return "u"
}