<?php

$queens = [ ];

$solutions = placeQueens( $queens );

foreach ( $solutions as $i => $solution ) {
	print ($i + 1 ) . ". " . implode( ' ', $solution ) . "\n";
}

/**
 * $fixedQueens is a one-dimensional array, where the key is the row index
 * and the value is the column index of the queen in that row. The number of
 * elements in the array is the number of fixed rows; the number of elements
 * plus one is the row index currently under study.
 *
 * The return value is a list of all solutions for the given preconditions, a 2-d
 * array. The key is arbitrary and the value is the queen positions for the rows
 * in the lower partition of the board. So each element of this array will have 8-n
 * elements where n is the number of elements in $fixedQueens.
 */
function placeQueens( $fixedQueens ) {
	// The row index under study, for convenience.
	$row = count( $fixedQueens );
	// Calculate all positions where the queen could be placed in the current row,
	// given $fixedQueens. This is an 8-element boolean valued array.
	$valid = getValidPositions( $fixedQueens );
	$solutions = [];
	for ( $col = 0; $col < 8; $col++ ) {
		if ( !$valid[$col] ) {
			continue;
		}
		if ( $row === 7 ) {
			// The terminating case of the recursion, all valid positions correspond to solutions
			$solutions[] = [ $col ];
		} else {
			// Deeply copy the queen positions and add the proposed position to the array
			$newQueens = $fixedQueens;
			$newQueens[] = $col;
			// Get the solutions for this proposed position
			$newSolutions = placeQueens( $newQueens );
			// For each of the solutions for the proposed position, add an element
			// to $solutions. If there are no solutions to the lower partition,
			// no solutions will be added to $solutions.
			foreach ( $newSolutions as $solution ) {
				$solutions[] = array_merge( [ $col ], $solution );
			}
		}
	}
	return $solutions;
}

/**
 * $queens is the positions of the fixed queens
 *
 * Returns an 8-element boolean valued array, false for invalid, true for valid.
 */
function getValidPositions( $queens ) {
	$nextRow = count( $queens );
	$valid = array_fill( 0, 8, true );
	// Consider each of the fixed queens in turn
	foreach ( $queens as $row => $col ) {
		$delta = $nextRow - $row;
		// Mark the position vertically below the queen as invalid
		$valid[$col] = false;
		// Compute the column indexes of the diagonals
		$plusDiag = $col + $delta;
		$minusDiag = $col - $delta;
		// Mark them invalid if they are on the board
		if ( $plusDiag < 8 ) {
			$valid[$plusDiag] = false;
		}
		if ( $minusDiag >= 0 ) {
			$valid[$minusDiag] = false;
		}
	}
	return $valid;
}
