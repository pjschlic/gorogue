package main

import (
  "fmt"
  termbox "github.com/nsf/termbox-go"
  "time"
  "os"
)

func log(s string) {
  fmt.Print(s)
}

type Pos struct {
  x, y int
}

func (p Pos) X() int {
  return p.x
}

func (p *Pos) SetX(x int) {
  p.x = x
}

func (p Pos) Y() int {
  return p.y
}

func (p *Pos) SetY(y int) {
  p.y = y
}

func (p *Pos) SetPos(x, y int) {
  p.x = x
  p.y = y
}

type Object struct {
  Pos
  cell termbox.Cell
}

func (o *Object) SetCell(c termbox.Cell) {
  o.cell = c
}

func (o Object) Cell() termbox.Cell {
  return o.cell
}

type Creature struct {
  Object
  health, mana int
  attribute [4]int
}

// Attribute enum:
type Attribute int
const (
    STR Attribute = iota
    DEX
    VIT
    MAG
)

func (c *Creature) SetHealth(h int) {
  c.health = h
}

func (c *Creature) SetMana(m int) {
  c.mana = m
}

func (c *Creature) SetAttribute(a Attribute, val int) {
  c.attribute[a] = val
}

func (c Creature) Mana() int {
  return c.mana
}

func (c Creature) Health() int {
  return c.health
}

type Player struct {
  Creature
}

func NewPlayer() *Player {
  p := Player{}
  p.SetPos(32, 32)
  p.SetCell(termbox.Cell{'@',termbox.ColorWhite, termbox.ColorBlack})
  // p.SetHealth(10)
  // p.SetMana(10)
  p.SetAttribute(STR, 10)
  return &p
}

func (o Object) Draw() {
  termbox.SetCell(o.x, o.y, o.cell.Ch, o.cell.Fg, o.cell.Bg)
}

func main() {
  termbox.Init()
  termbox.HideCursor()
  var p Player
  p = *NewPlayer()
  p.Draw()
  termbox.Flush()
  <-time.After(2000 * time.Millisecond) //++ TODO: find out wether .After() or .Sleep() is better performance-wise
  termbox.Close()
  os.Exit(0)
}

