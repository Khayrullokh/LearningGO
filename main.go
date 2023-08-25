package main

import "fmt"

func main() {
	var a float64
	var b float64
	var c float64
	var d float64

	fmt.Print("Input floats a , b, c, d: ")
	fmt.Scanln(&a, &b, &c, &d)
	fmt.Printf("\na = %f, b = %f, c = %f, d = %f\n", a, b, c, d)

	if a > b {
		if a > c {
			if a > d {
				if b > c {
					if b > d {
						fmt.Println("a > b > c > d")
					} else if d > c {
						fmt.Println("a > b > d > c")
					}
				} else if c > b {
					if c > d {
						fmt.Println("a > c > b > d")
					} else if d > b {
						fmt.Println("a > c > d > b")
					}
				} else if d > b {
					if d > c {
						fmt.Println("a > d > b > c")
					} else if c > b {
						fmt.Println("a > d > c > b")
					}
				} else if b > a {
					if b > c {
						if b > d {
							if a > c {
								if a > d {
									fmt.Println("b > a > c > d")
								} else if d > c {
									fmt.Println("b > a > d > c")
								}
							} else if c > a {
								if c > d {
									if a > d {
										fmt.Println("b > c > a > d")
									} else if d > a {
										fmt.Println("b > c > d > a")
									}
								}
							} else if d > a {
								if d > c {
									if a > c {
										fmt.Println("b > d > a > c")
									} else if c > a {
										fmt.Println("b > d > c > a")
									}
								}
							}
						}
					}
				} else if c > a {
					if c > b {
						if c > d {
							if a > b {
								if b > d {
									fmt.Println("c > a > b > d")
								} else if d > b {
									fmt.Println("c > a > d > b")
								}
							} else if b > a {
								if b > d {
									if a > d {
										fmt.Println("c > b > a > d")
									} else if d > a {
										fmt.Println("c > b > d > a")
									}
								}
							} else if d > a {
								if d > b {
									if a > b {
										fmt.Println("c > d > a > b")
									} else if b > a {
										fmt.Println("c > d > b > a")
									}
								}
							}
						}
					}
				} else if d > a {
					if d > c {
						if d > b {
							if a > c {
								if a > b {
									fmt.Println("d > a > c > b")
								} else if b > c {
									fmt.Println("d > a > b > c")
								}
							} else if c > a {
								if c > b {
									if a > b {
										fmt.Println("d > c > a > b")
									} else if b > a {
										fmt.Println("d > c > b > a")
									}
								}
							} else if b > a {
								if b > c {
									if a > c {
										fmt.Println("d > b > a > c")
									} else if c > a {
										fmt.Println("d > b > c > a")
									}
								}
							}
						}
					}
				}
			}
		}

	}
}
