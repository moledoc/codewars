package bouncingballs

func BouncingBall(h, bounce, window float64) int {
	// Check parameter conditions
	if !(h > 0) || !(bounce > 0 && bounce < 1) || !(h > window) {
		return -1
	}
	// Ball will always be seen once (when the ball is falling, since h > window)
	// Will init the counter `ballSeen` from -1, because the first iteration of the loop always happens (h>window),
	// meaning if we only see ball falling, then the end result is -1 +2 = 1.
	var ballSeen int = -1
	for ; h > window; h *= bounce {
		// When ball bounces over the window, mother sees the ball twice:
		// once going up, once going down.
		ballSeen += 2
	}
	return ballSeen
}
