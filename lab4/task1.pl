act(  
    go(Start, End),
    [current(Start), connection(Start, End), on(floor)],
    [current(Start)],
    [current(End)]
).

act( 
    push(Box, Start, End),
    [box(Box), on(floor), current(Start), currentBox(Box, Start), connection(Start, End), connection(Light, Start), light(Light, on)],
    [current(Start), currentBox(Box, Start)],
    [current(End), currentBox(Box, End)]
).

act(  
    climbUp(Box),
    [on(floor), box(Box), current(X), currentBox(Box, X)],
    [on(floor)],
    [on(Box)]
).

act(
    climbDown(Box),
    [on(Box), box(Box), current(X), currentBox(Box, X)],
    [on(Box)],
    [on(floor)]
).

act(  
    turnOn(Light),
    [on(X), current(Light), currentBox(X, Light), light(Light, off)],
    [light(Light, off)],
    [light(Light, on)]
).

act(  
    turnOff(Light),
    [on(X), current(Light), currentBox(X, Light), light(Light, on)],
    [light(Light, on)],
    [light(Light, off)]
).

% 1. to move Shakey from room3 to room1
%goal_state([ current(room1) ]).

% 2. to switch off the light in room1
goal_state([ light(switch1, off) ]).

% 3. to get box2 into room2 
%goal_state([ currentBox(box2, room2) ]).

initial_state([
    current(room3),
    on(floor),

    box(box1),
    box(box2),
    box(box3),
    box(box4),

    % current
    currentBox(box1, room1),
    currentBox(box2, room1),
    currentBox(box3, room1),
    currentBox(box4, room1),

    % room -> corridor
    connection(room1, corridor),
    connection(room2, corridor),
    connection(room3, corridor),
    connection(room4, corridor),
    connection(corridor, room1),
    connection(corridor, room2),
    connection(corridor, room3),
    connection(corridor, room4),

    % room -> light
    connection(room1, switch1),
    connection(room2, switch2),
    connection(room3, switch3),
    connection(room4, switch4),
    connection(switch1, room1),
    connection(switch2, room2),
    connection(switch3, room3),
    connection(switch4, room4),
    connection(xxx, corridor),

    light(switch1, on),
    light(switch2, off),
    light(switch3, off),
    light(switch4, on),
    light(xxx, on)
]).
