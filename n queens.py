
class Queen:
    def __init__(self, pos, n):
        self.pos = pos
        self.board_size = n
        self.blocking_positions = self.get_blocking_positions(pos)

    def get_blocking_positions(self, pos):
        bpos = []
        for i in range(self.board_size):
            bpos.append((self.pos[0], i))
            bpos.append((i, self.pos[1]))
        bpos += self.get_diag_bpos(pos)
        return bpos

    def get_diag_bpos(self, pos):
        bpos = []
        row = pos[0]
        col = pos[1]
        for i in range(self.board_size):
            hdist = row-i
            if col + hdist <= self.board_size:
                bpos.append((i, col+hdist))
            if col - hdist >= 0:
                bpos.append((i, col-hdist))
        print(bpos)
        return bpos
        
            

def create_array(n):
    erow = [0 for i in range(n)]
    array = [erow for i in range(n)]
    return array


def add_queen(n):
    queens.append(Queen((0,0), n))
    board[0][0] = 1
    print(board)


n = int(input('Board Size: '))
queens = []
board = create_array(n)
add_queen(n)
print(board)
for a in board:
    print(a)
