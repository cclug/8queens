<?php

$queens = [ ];

$solutions = placeQueens( $queens );

foreach ( $solutions as $i => $solution ) {
	print ($i + 1 ) . ". " . implode( ' ', $solution ) . "\n";
}

function placeQueens( $fixedQueens ) {
	$row = count( $fixedQueens );
	$valid = getValidPositions( $fixedQueens );
	$solutions = [];
	for ( $col = 0; $col < 8; $col++ ) {
		if ( !$valid[$col] ) {
			continue;
		}
		if ( $row === 7 ) {
			$solutions[] = [ $col ];
		} else {
			$newQueens = $fixedQueens;
			$newQueens[] = $col;
			$newSolutions = placeQueens( $newQueens );
			foreach ( $newSolutions as $solution ) {
				$solutions[] = array_merge( [ $col ], $solution );
			}
		}
	}
	return $solutions;
}

function getValidPositions( $queens ) {
	$nextRow = count( $queens );
	$valid = array_fill( 0, 8, true );
	foreach ( $queens as $row => $col ) {
		$delta = $nextRow - $row;
		$valid[$col] = false;
		$plusDiag = $col + $delta;
		$minusDiag = $col - $delta;
		if ( $plusDiag < 8 ) {
			$valid[$plusDiag] = false;
		}
		if ( $minusDiag >= 0 ) {
			$valid[$minusDiag] = false;
		}
	}
	return $valid;
}
