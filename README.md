# pokedexcli

## ENG

This is my first Go (working) program where I implement HTTP client logic!
This REPL (Read-Eval-Print-Loop) is a Go program uses the api from [text](https://pokeapi.co).
It is a command line interface that accepts some basic commands to make requests using the 
aforementioned api, and if the request is succesful, the JSON data response received is parsed in order
to be printed to the user.

The program has an implementation of a cache system to hold data from requests (for faster availability) for a 
configurable  interval of time. Every time the interval of time passes, a tick calls a method (reapLoop) to clean
from the cache from all data with a permanence longer than said interval.

Some of the implemented commands are:
- help : Displays all available commands.
- exit : Exits the program.
- map : Displays 20 location areas.
- explore : Displays list of Pokemon inhabiting an especific area.
- catch : Throws a pokeball to a given Pokemon in an attempt to catch it.
- inspect : Displays stats from a Pokemon.
- pokedex : Displays the names of all Pokemon caught by the user.

_Some improvements to the program are due to be implemented in the future_

## ESP

Este es mi primer programa en Go (funcional) donde implemento lógica de clientes HTTP.
Este REPL (Read-Evaluate-Print-Loop) es un programa de Go que usa la api de [text](https://pokeapi.co).
Es una interfaz de líneas de comando que acepta algunos comandos basicos para hacer pedidos usando dicha api,
y si el pedido es exitoso, se analiza la respuesta de JSON recibida para ser impresa al usuario.

El programa tiene la implementación de un sistema de caché para retener datos de los pedidos (para una disponibilidad mas veloz)
durante un intervalo de tiempo configurable. Cada vez que transcurre el tiempo del intervalo, un tick llama a un método (reapLoop)
para limpiar el caché de los datos que hayan tenido una permanencia mayor a la de dicho intervalo.

Algunos de los comandos implementados son:
- help : Despliega información sobre todos los comandos disponibles.
- exit : Sale del programa.
- map : Despliega 20 áreas de ubicación.
- explore : Muestra los Pokemon que se encuentran en una dada área.
- catch : Lanza una Pokebola a un Pokemon designado para intentar atraparlo.
- inspect : Despliega los datos sobre un dado Pokemon.
- pokedex : Despliega una lista de todos los Pokemon atrapados por el usuario.

_Algunas mejoras están sujetas a ser implementadas en el futuro_