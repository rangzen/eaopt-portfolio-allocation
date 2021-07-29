# eaopt-portfolio-allocation

Use the GA library eaopt for portfolio allocation.

# Usage

```shell
git clone https://github.com/rangzen/eaopt-portfolio-allocation
cd eaopt-portfolio-allocation
go build
```

# Specificities

## Only buy

If you set the number of shares that you already have in the owned field, this will
be a minimum in the number of shares. The algorithm does not propose sells. 

# JSON Data

## Targets

Allocations targets are percentages. The names are the same that you should find in
allocations’ shares.

## Shares

The curr_ratio is the ratio of the general currency. For example, if all the prices
are in Euro and some shares are in US dollars, you can set the curr_ratio of
those shares at .66.

# Links

* https://github.com/MaxHalford/eaopt