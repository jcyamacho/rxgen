package funcutil

import "log"

func Fatal(fn func() error) func() {
	return func() {
		if err := fn(); err != nil {
			log.Fatal(err)
		}
	}
}

func Fatal1[P1 any](fn func(P1) error) func(P1) {
	return func(p1 P1) {
		if err := fn(p1); err != nil {
			log.Fatal(err)
		}
	}
}

func Fatal2[P1, P2 any](fn func(P1, P2) error) func(P1, P2) {
	return func(p1 P1, p2 P2) {
		if err := fn(p1, p2); err != nil {
			log.Fatal(err)
		}
	}
}

func Fatal3[P1, P2, P3 any](fn func(P1, P2, P3) error) func(P1, P2, P3) {
	return func(p1 P1, p2 P2, p3 P3) {
		if err := fn(p1, p2, p3); err != nil {
			log.Fatal(err)
		}
	}
}
