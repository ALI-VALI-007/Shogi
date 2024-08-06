package main
import "fmt"

//Player Stuff
type player struct{
	reserves map[string]int //rook, bishop, gold, silver, knight, lance, pawn
	reigning bool //Reigning means moves first and also dif king (white for chess basically)
	moves Queue //Each player will have their own moves queue so that they can be emptied if their king is in check
	placedPieces []piece
}
func (p *player) add(newPiece string){
	reserves[newPiece]+=1
}
func (p *player) remove(pieceName string) bool{
	reserves[newPiece]-=1
}
//This is the queue we are gonna use for premoves for each player
type Queue struct{
	Elements []move
}
func (q *Queue) add(elem newMove){
	q.Elements.add(newMove)
}
func (q *Queue) pop() move{
	if len(q.Elements)==0{
		return
	}
	result=q.Elements[0]
	q.Elements=q.Elements[1:]
	return result
}
func (q *Queue) empty(){
	q.Elements=nil
}

//Pieces
//We can use the general piece struct for them to move around, we need the specific x and y range of movements however for each piece
type piece struct{
	isSelected bool
	coordinates [2]int
	isPromoted bool
	name string
	moveSet [][2]int
	isReigning bool
}

func move(attacker *player,b *board,p *piece,x int, y int){
	captured:=b.grid[x][y]
	b.grid[x][y]=p
	p.coordinates[0]=x
	p.coordinates[1]=y
	if captured!=nil{
		attacker.add(captured)
	}
}

type king struct{
	isChecked bool
	isMated bool
	base piece
}

//Environment
//Boards are 9*9
type board struct{
	grid [9][9]boardNode
	placedPieces []piece
}

type boardNode struct{
	Value piece
	isHighlighted bool
	//Adjacent []*boardNode
}

/*func capture(attacker *player, p *piece){
	player.add(p)
}*/
/*func canCapture(attacker *piece,captured *piece){
	if attacker.isReigning!=captured.isReigning{
		return true
	}
	return false
}*/

func canDeploy(attacker *player, p string,b *board, x int, y int) bool{
	newPiece:=attacker.remove(p)
	cur:=&b.grid[newPiece.coordinates[0]][newPiece.coordinates[1]]
	attacker.add(p)
	if cur.Value!=nil{
		return false
	} else{
		if p=="pawn"{
			if attacker.isReigning && b.grid[x][y+1].Value.name == "king" && b.grid[x][y+1].Value.isReigning!=attacker.isReigning {
				//cur.Value=newPiece
				return true
			} else if !attacker.isReigning && b.grid[x][y-1].Value.name == "king" && b.grid[x][y-1].Value.isReigning!=attacker.isReigning {
				//cur.Value=newPiece
				return true
			} else{
				return false
			}
		}
		else{
			//cur.Value=newPiece
			return true
		}
	}
	return false
}
func deploy(attacker *player, p string, b *board, x int, y int){
	if canDeploy(attacker,p,b,x,y){
		attacker.remove(p)
		newPiece:=createPiece(p,attacker)
		newPiece.isReigning=attacker.isReigning
		b.coordinates[x][y]=newPiece
	}
}
func createPiece(p string,attacker *player) piece{
	switch p{
		case "rook":
			return rook()
		case "bishop":
			return bishop()
		case "gold":
			return gold()
		case "silver":
			return silver()
		case "knight":
			return knight()
		case "lance":
			return lance()
		case "pawn":
			return pawn()
	}
	return
}
func place(b *board, attacker *player, p *piece, x int, y int){
	b.coordinates[p.coordinates[0]][p.coordinates[1]]=nil
	p.move(x,y)
	b.coordinates[p.coordinates[0]][p.coordinates[1]]=p
}
/*type piece struct{
	isSelected bool
	coordinates [2]int
	isPromoted bool
	name string
	moveSet [][2]int
}*/
//rook, bishop, gold, silver, knight, lance, pawn
func rook() piece{
	return piece{
		isSelected: false,
		isPromoted: false,
		name: "rook"
	}
}
func potentialMovesRook(p *piece, b *board){
	x:=p.coordinates[0]
	y:=p.coordinates[1]
	for i:=x+1; i<9; i++{
		cur:=b.coordinates[i][y]
		if cur.Value==nil || cur.Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(i,y))}
		if cur.Value!=nil{
			break
		}
	}
	for i:=x-1; i>=0; i++{
		cur:=b.coordinates[i][y]
		if cur.Value==nil || cur.Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(i,y))}
		if cur.Value!=nil{
			break
		}
	}
	for i:=y+1; i<9; i++{
		cur:=b.coordinates[x][i]
		if cur.Value==nil || cur.Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x,i))}
		if cur.Value!=nil{
			break
		}
	}
	for i:=y-1; i>=0; i++{
		cur:=b.coordinates[x][i]
		if cur.Value==nil || cur.Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x,i))}
		if cur.Value!=nil{
			break
		}
	}

	if p.isPromoted {
		if x+1<9 {
			if y+1<9{
				if b.coordinates[x+1][y+1].Value==nil || b.coordinates[x+1][y+1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x+1,y+1))}
			}
			if y-1>=0{
				if b.coordinates[x+1][y-1].Value==nil || b.coordinates[x+1][y-1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x+1,y-1))}
			}
		}
		if x-1>=0{
			if y+1<9{
				if b.coordinates[x-1][y+1].Value==nil || b.coordinates[x-1][y+1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x-1,y+1))}
			}
			if y-1>=0{
				if b.coordinates[x-1][y-1].Value==nil || b.coordinates[x-1][y-1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x-1,y-1))}
			}
		}
	}
}
func bishop() piece{
	return piece{
		isSelected: false,
		isPromoted: false,
		name: "bishop"
	}
}
func potentialMovesBishop(b *board,p *piece){
	x:=p.coordinates[0]
	y:=p.coordinates[1]
	for i:=1; i+x<9 && i+y<9 ; i++{
		if b.coordinates[x+i][y+i].Value==nil || b.coordinates[x+i][y+i].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet, (i+x,i+y))}
		if b.coordinates[i+x][i+y].Value!=nil {
			break
		}
	}for i:=1; i+x<9 && y-i>=0 ; i++{
		if b.coordinates[x+i][y-i].Value==nil || b.coordinates[x+i][y-i].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet, (i+x,y-i))}
		if b.coordinates[i+x][y-i].Value!=nil {
			break
		}
	}for i:=1; x-i>=0 && y+i<9 ; i++{
		if b.coordinates[x-i][y+i].Value==nil || b.coordinates[x-i][y+i].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet, (x-i,y+i))}
		if b.coordinates[x-i][y+i].Value!=nil {
			break
		}
	}for i:=1; x-i>=0 && y-i>=0 ; i++{
		if b.coordinates[x-i][y-i].Value==nil || b.coordinates[x-i][y-i].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet, (x-i,y-i))}
		if b.coordinates[x-i][y-i].Value!=nil {
			break
		}
	}
	if p.isPromoted{
		if x+1<9 {
			if b.coordinates[x+1][y].Value==nil || b.coordinates[x+1][y].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x+1,y))}
		}
		if y-1>=0{
			if b.coordinates[x][y-1].Value==nil || b.coordinates[x][y-1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x,y-1))}
		}
		if y+1<9{
			if b.coordinates[x][y+1].Value==nil || b.coordinates[x][y+1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x,y+1))}
		}
		if x-1>=0{
			if b.coordinates[x-1][y].Value==nil || b.coordinates[x-1][y].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x-1,y))}
		}
	}
}
func gold() piece{
	return piece{
		isSelected: false,
		isPromoted: false,
		name: "gold"
	}
}
func potentialMovesGold(p *piece){
	x:=p.coordinates[0]
	y:=p.coordinates[1]
	if p.isReigning {
		if y+1<9{
			p.moveSet=append(p.moveSet,(x,y+1))
			if x+1<9{
				if b.coordinates[x+1][y+1].Value==nil || b.coordinates[x+1][y+1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x+1,y+1))}
			}
			if x-1>=0{
				if b.coordinates[x-1][y+1].Value==nil || b.coordinates[x-1][y+1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x-1,y+1))}
			}
		}
		if x+1<9{
			if b.coordinates[x+1][y].Value==nil || b.coordinates[x+1][y].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x+1,y))}
		}
		if x-1>=0{
			if b.coordinates[x-1][y].Value==nil || b.coordinates[x-1][y].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x-1,y))}
		}
		if y-1>=0{
			if b.coordinates[x][y-1].Value==nil || b.coordinates[x][y-1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x,y-1))}
		}
	} else{
		if y-1>=0{
			p.moveSet=append(p.moveSet,(x,y-1))
			if x+1<9{
				if b.coordinates[x+1][y-1].Value==nil || b.coordinates[x+1][y-1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x+1,y-1))}
			}
			if x-1>=0{
				if b.coordinates[x-1][y-1].Value==nil || b.coordinates[x-1][y-1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x-1,y-1))}
			}
		}
		if x+1<9{
			if b.coordinates[x+1][y].Value==nil || b.coordinates[x+1][y].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x+1,y))}
		}
		if x-1>=0{
			if b.coordinates[x-1][y].Value==nil || b.coordinates[x-1][y].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x-1,y))}
		}
		if y+1<9{
			if b.coordinates[x][y+1].Value==nil || b.coordinates[x][y+1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x,y+1))}
		}
	}
}
func silver() piece{
	return piece{
		isSelected: false,
		isPromoted: false,
		name: "silver"
	}
}
func potentialMovesSilver(p *piece){
	if p.isPromoted{
		potentialMovesGold(p)
	} else{
		x:=p.coordinates[0]
		y:=p.coordinates[1]
		if p.isReigning {
			if y+1<9{
				if b.coordinates[x][y+1].Value==nil || b.coordinates[x][y+1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x,y+1))}
				if x+1<9{
					if b.coordinates[x+1][y+1].Value==nil || b.coordinates[x+1][y+1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x+1,y+1))}
				}
				if x-1>=0{
					if b.coordinates[x-1][y+1].Value==nil || b.coordinates[x-1][y+1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x-1,y+1))}
				}
			}
			if y-1>=0{
				if x-1>=0{
					if b.coordinates[x-1][y-1].Value==nil || b.coordinates[x-1][y-1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x-1,y-1))}
				}
				if x+1<9{
					if b.coordinates[x+1][y-1].Value==nil || b.coordinates[x+1][y-1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x+1,y-1))}
				}
			}
		} else{
			if y-1>=0{
				if b.coordinates[x][y-1].Value==nil || b.coordinates[x][y-1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x,y-1))}
				if x+1<9{
					if b.coordinates[x+1][y-1].Value==nil || b.coordinates[x+1][y+1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x+1,y-1))}
				}
				if x-1>=0{
					if b.coordinates[x-1][y-1].Value==nil || b.coordinates[x-1][y+1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x-1,y-1))}
				}
			}
			if y+1<9{
				if x-1>=0{
					if b.coordinates[x-1][y+1].Value==nil || b.coordinates[x-1][y+1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x-1,y+1))}
				}
				if x+1<9{
					if b.coordinates[x+1][y+1].Value==nil || b.coordinates[x+1][y+1].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x+1,y+1))}
				}
			}
		}

	}
}
	return piece{
		isSelected: false,
func knight() piece{
		isPromoted: false,
		name: "knight"
	}
}
func potentialMovesKnight(b *board,p *piece){
	if p.isPromoted{
		potentialMovesGold(p)
	} else{
		x:=p.coordinates[0]
		y:=p.coordinates[1]
		if p.isReigning {
			if y+2<9{
				if x+1<9{
					if b.coordinates[x+1][y+2].Value==nil || b.coordinates[x+1][y+2].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x+1,y+2))}
				}
				if x-1>=0{
					if b.coordinates[x-1][y+2].Value==nil || b.coordinates[x-1][y+2].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x-1,y+2))}
				}
			}
		} else{
			if y-2>=0{
				if x+1<9{
					if b.coordinates[x+1][y-2].Value==nil || b.coordinates[x+1][y-2].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x+1,y-2))}
				}
				if x-1>=0{
					if b.coordinates[x-1][y-2].Value==nil || b.coordinates[x-1][y-2].Value.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x-1,y-2))}
				}
			}
		}
	}
}
func lance() piece{
	return piece{
		isSelected: false,
		isPromoted: false,
		name: "lance"
	}
}
func potentialMovesLance(b *board,p *piece) {
	if p.isPromoted{
		potentialMovesGold(p)
	} else{
		x:=p.coordinates[0]
		y:=p.coordinates[1]
		if p.isReigning {
			for i:=y;i<9;i++{
				cur:=b.coordinates.Value
				if cur==nil || cur.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x,i))}
				if cur!=nil{
					break
				}
			}
		}else if !p.isReigning {
			for i:=y;i>=0;i--{
				cur:=b.coordinates.Value
				if cur==nil || cur.isReigning!=p.isReigning {p.moveSet=append(p.moveSet,(x,i))}
				if cur!=nil{
					break
				}
			}
		}
	
	}
}
func pawn() piece{
	return piece{
		isSelected: false,
		isPromoted: false,
		name: "pawn"
	}
}
func potentialMovesPawn(b *board,p *piece){
	if p.isPromoted{
		potentialMovesGold(p)
	} else{
		x:=p.coordinates[0]
		y:=p.coordinates[1]
		if p.isReigning && (b.coordinates[x][y+1]==nil || p.isReigning!=b.coordinates[x][y+1].Value.isReigning) {p.moveSet=append(p.moveSet,(x,y+1))}
		else if !p.isReigning && (b.coordinates[x][y-1]==nil || p.isReigning!=b.coordinates[x][y-1].Value.isReigning) {p.moveSet=append(p.moveSet,(x,y-1))}
	}
}
func legalMovesChecker(p *piece, k *king,b *board) [][2]int{
	curMoveSet:=p.moveSet
	x:=p.coordinates[0]
	y:=p.coordinates[1]
	for i:=0;i<curMoveSet.length();i++{
		curMove:=curMoveSet[i]
		xP:=curMove[0]
		yP:=curMove[1]
		p.move(xP,yP)
		//Check func for if king is in check (board func)
		b.isChecked(k)
		if k.isChecked {
			curMoveSet=curMoveSet[:i]
			curMoveSet=append(curMoveSet[i+1:])
		}
	}
	return curMoveSet
}
func (b *board)isChecked(k *king){
	x:=k.base.coordinates[0]
	y:=k.base.coordinates[1]
	pieces:=b.placedPieces
	for i:=0;i<pieces.length();i++{
		curPiece:=pieces[i]
		if curPiece.isReigning!=k.isReigning {
			for j:=0;j<curPiece.moveSet.length();i++{
				if curPiece.moveSet[j][0]==x && curPiece.moveSet[j][1]==y {
					k.isChecked=true
				}
			}
		}
	}
}
func (b *board)hasCheckMate(k *king)bool{
	allPieces:=b.placedPieces
	var sameSidePiece piece
	for i:=0;i<allPieces.length();i++{
		if allPieces[i].isReigning==k.isReiging{
			sameSidePiece=append(allPieces[i])
		}
	}
	for i:=0;i>sameSidePiece.length();i++{
		if legalMovesChecker(sameSidePiece[i],k,b).length()>0{
			return false
		}
	}
	return true
}
func isDraw(reigning *king,challenger *king,b *board) bool{
	if b.hasCheckMate(reigning) && b.hasCheckMate(challenger){
		return true
	}
	return false
}
func createBoard(){
	b:=new board{

	}
	var allPieces []piece
	for i:=0;i<9;i++{
		reigningPawn:=pawn()
		reigningPawn.isReigning=true
		challengerPawn:=pawn()
		challengerPawn.isReigning=false
		b.grid[i][3]=reigningPawn
		b.grid[i][6]=challengerPawn
	}
	reigningBishop:=bishop()
	reigningBishop.isReigning=true
	b.grid[1][1]=reigningBishop
	challengerBishop:=bishop()
	challengerBishop.isReigning=false
	b.grid[7][7]=challengerBishop

	reigningRook:=rook()
	reigningRook.isReigning=true
	b.grid[7][1]=reigningRook
	challengerRook:=rook()
	challengerRook.isReigning=false
	b.grid[1][7]=challengerRook

	reigningLance:=lance()
	reigningLance.isReigning=true
	b.grid[0][0]=reigningLance
	reigningLance:=lance()
	reigningLance.isReigning=true
	b.grid[8][0]=reigningLance
	reigningLance:=lance()
	reigningLance.isReigning=false
	b.grid[0][8]=reigningLance
	reigningLance:=lance()
	reigningLance.isReigning=false
	b.grid[8][8]=reigningLance


}
