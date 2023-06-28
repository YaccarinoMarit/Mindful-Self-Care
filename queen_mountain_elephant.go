package msCare

import "fmt"

//MindfulSelfCare contains method for taking care of one's self
type MindfulSelfCare struct {
    SelfAwareness int
}

//BeAware launches a series of activities to become more aware of yourself
func (m *MindfulSelfCare) BeAware() {
    for i := 0; i < m.SelfAwareness; i++ {
        //do something to increase your awareness
    }
    fmt.Println("You have increased your self-awareness!")
}

//Meditate initiates a guided meditation session
func (m *MindfulSelfCare) Meditate() {
    //play soothing music
    //offer guided meditation
    fmt.Println("You have completed your meditation session!")
}

//Exercise initiates an exercise routine tailored to user's needs
func (m *MindfulSelfCare) Exercise() {
    //set up routine
    fmt.Println("You have completed your exercise session!")
}

//TakeABreak is designed to take a break from any taxing activities
func (m *MindfulSelfCare) TakeABreak() {
    //take a few minutes off
    fmt.Println("You have taken a break! You have reset your focus.")
}

//EatWell is designed to promote healthy eating habits
func (m *MindfulSelfCare) EatWell() {
    //cook a healthy meal
    //find healthy snacks
    fmt.Println("You have eaten a healthy meal!")
}

//GoOutside is designed to get user outside and enjoy nature
func (m *MindfulSelfCare) GoOutside() {
    //go for a short walk
    //find a calming spot outdoors
    fmt.Println("You have refreshed your senses with nature!")
}

//CatchUp is designed to stay in touch with friends and family
func (m *MindfulSelfCare) CatchUp() {
    //call a friend
    //write someone a letter
    fmt.Println("You have caught up with your loved ones!")
}