package keeper

import (
	"github.com/BenWolfaardt/Checkers/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
