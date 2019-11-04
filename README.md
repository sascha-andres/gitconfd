# Git conf.d directory lib

You provide the type of the hook and the bnase directory. This lib executes all
scripts within <hook-type>.d

Eg commit.d/

## Directory layout

Within the directory there can be platform specific and agnostic scripts. The platform specific files must be 
located within a sub directory named like the GOOS:

    .commit-msg.d/windows/
    .commit-msg.d/linux/

Platform agnostic scripts are located within a special directory name `agnostic`

## History

|Version|Description|
|---|---|
|0.1.0|Initial implementation|
