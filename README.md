# pokedexProject
<pre>
   __ ___         _            __     ___      ___
   | '_  \       | |           | |    \  \    /  /
   | |_) |  ___  | | _____  ___| |___  \  \  /  /
   | .___/ / _ \ | |/ / _ \/  _  | _ \  \  \/  /
   | |    | (_) ||   <| __/| (_) | __/  /  /\  \
   | |     \___/ |_|\_\___/\_____|___/ /  /  \  \
   |_|                                /__/    \__\
</pre>
The Pokedex generates ascii art of all Pokemon that are searched. If you're on MacOS, there's also some neat goofy voice action.

I used the PokéAPI for looking up Pokemon in this project. 
It's pretty neat so if you're interested, check that out over at https://pokeapi.co/

The ascii art code is a reimplementation of the "convert" package developed by Qeesung as a part of "Image2Ascii" over at https://github.com/qeesung/image2ascii

![alt text](raichu.png)

UPDATE:

The GoFx microservice sample was removed due to issues with ascii conversion and colorization.
Refactored to better organize tasks into packages.

Additional work:
- Add terminal frames to separate areas of Pokedex content. This will better organize the output.
