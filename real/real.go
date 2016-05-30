package real

type Reality struct {
  Emptity bool
  May bool
  Be int64
}

func (r *Reality) Play() {
  go func (real *Reality) {
    if real.Emptity {
      println(`you is nothingness`)
    } else if real.May {
      println(`you is illusion`)
    } else if {
      println(`you is being`)
    } else {
      println(`you is ...`)
    }
  }(r)
}
