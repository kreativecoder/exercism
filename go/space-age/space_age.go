package space

const earth = 31557600
const mercury = 0.2408467 * earth
const venus = 0.61519726 * earth
const mars = 1.8808158 * earth
const jupiter = 11.862615 * earth
const saturn = 29.447498 * earth
const uranus = 84.016846 * earth
const neptune = 164.79132 * earth

//Age : accepts age in seconds from planet and returns a persons age on earth
func Age(age float64, planet string) float64 {
	switch planet {
	case "Mercury":
		return age / mercury
	case "Venus":
		return age / venus
	case "Mars":
		return age / mars
	case "Jupiter":
		return age / jupiter
	case "Saturn":
		return age / saturn
	case "Uranus":
		return age / uranus
	case "Neptune":
		return age / neptune
	default:
		return age / earth
	}
}
