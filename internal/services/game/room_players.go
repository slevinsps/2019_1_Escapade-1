package game

import "fmt"

// RecoverPlayer call it in lobby.join if player disconnected
func (room *Room) RecoverPlayer(i int, newConn *Connection) {

	oldConn := room.Players.Connections[i]

	if !oldConn.disconnected {
		oldConn.Kill("Another connection found", true)
		//oldConn.SendInformation([]byte("Another connection found"))
	}

	// add connection as player
	room.MakePlayer(newConn)

	//room.Players.Connections[i] = newConn

	room.addAction(newConn, ActionReconnect)
	room.sendHistory(AllExceptThat(newConn))

	return
}

// RecoverObserver recover connection as observer
func (room *Room) RecoverObserver(oldConn *Connection, newConn *Connection) {

	if !oldConn.disconnected {
		oldConn.Kill("Another connection found", true)
		//oldConn.SendInformation([]byte("Another connection found"))
	}

	room.MakeObserver(newConn)

	room.addAction(newConn, ActionReconnect)
	room.sendHistory(AllExceptThat(newConn))

	return
}

// observe try to connect user as observer
func (room *Room) addObserver(conn *Connection) bool {
	// if we havent a place
	if !room.Observers.enoughPlace() {
		conn.SendInformation([]byte("Room cant execute request "))
		return false
	}
	room.MakeObserver(conn)

	room.addAction(conn, ActionConnectAsObserver)

	room.sendObservers(AllExceptThat(conn))

	return true
}

// EnterPlayer handle player try to enter room
func (room *Room) addPlayer(conn *Connection) bool {
	// if room have already started
	// if room.Status != StatusPeopleFinding {
	// 	return false
	// }

	conn.debug("Room(" + room.Name + ") wanna connect you")

	// if room hasnt got places
	if !room.Players.enoughPlace() {
		conn.debug("Room(" + room.Name + ") hasnt any place")
		return false
	}

	room.MakePlayer(conn)

	room.addAction(conn, ActionConnectAsPlayer)
	room.sendPlayers(room.All)

	if !room.Players.enoughPlace() {
		room.startFlagPlacing()
	}

	return true
}

// MakePlayer mark connection as connected as Player
// add to players slice and set flag inRoom true
func (room *Room) MakePlayer(conn *Connection) {
	if room.Status != StatusPeopleFinding {
		room.lobby.waiterToPlayer(conn, room)
		conn.both = false
	} else {
		conn.both = true
	}
	room.Players.Add(conn)
	conn.PushToRoom(room)
}

// MakeObserver mark connection as connected as Observer
// add to observers slice and set flag inRoom true
func (room *Room) MakeObserver(conn *Connection) {
	if room.Status != StatusPeopleFinding {
		room.lobby.waiterToPlayer(conn, room)
		conn.both = false
	} else {
		conn.both = true
	}
	room.Observers.Add(conn)
	conn.PushToRoom(room)
}

func (room *Room) removeBeforeLaunch(conn *Connection) {
	room.Players.Remove(conn)
	fmt.Println("removing", len(room.Players.Connections))
	conn.debug("you went back to lobby")
	if room.Players.Empty() {
		room.Close()
		conn.debug("We closed room :С")
	}
}

func (room *Room) removeDuringGame(conn *Connection) {
	i := room.Players.Search(conn)
	if i >= 0 {
		room.GiveUp(conn)
		room.sendHistory(room.All)
		room.sendPlayers(room.All)
	} else {

		i := room.Observers.Search(conn)
		room.Observers.Remove(i)
		room.sendObservers(room.All)
	}
	if room.Players.Empty() {
		room.Close()
		conn.debug("We closed room :С")
	}
}

// removeFinishedGame
// func (room *Room) removeAfterLaunch(conn *Connection) {
// 	i := room.Players.Search(conn)
// 	if i >= 0 {
// 		room.TryClose()
// 		return
// 	}

// 	room.Observers.Remove(conn)
// 	room.sendObservers(room.all())
// 	room.TryClose()
// 	return
// }
