module "https://github.com/alejandrosanchezortiz/algblockchain"

require blockchain v0.0.0

replace blockchain v0.0.0 => ./blockchain

go 1.12
