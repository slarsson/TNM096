% STRIPS planner

plan :-
     initial_state(IS),
     goal_state(GS),
     solve(IS,GS,[],Plan),
     printPlan(Plan).


% if Goal is a subset of State, then return Plan
solve(State, Goal, Plan, Plan):- is_subset(Goal, State).

% otherwise, select next action and move to the next state
solve(State, Goal, Sofar, Plan):-
     act(Action, Precons, Delete, Add),
     is_subset(Precons, State),
     %writeNLNL(Sofar),
     \+ member(Action, Sofar),    % negation as failure
     delete_list(Delete, State, Remainder),
     add_list(Add, Remainder, NewState),
     solve(NewState, Goal, [Action|Sofar], Plan).



% AUXILIARY

% Check is first list is a subset of the second
is_subset([], _).
is_subset([H|T], Set):- member(H, Set), is_subset(T, Set).

printPlan(Plan) :-
     writeNL('--- One Plan ---'),
     printPlan2(Plan),
     writeNL('-----------------').

printPlan2([]).
printPlan2([H|T]):- printPlan2(T), writeNL(H).

% Remove all elements of 1st list from second to create third.
delete_list([], State, State) :- !.
delete_list([H|T], State, Newstate):-
           remove(H, State, Remainder), delete_list(T, Remainder, Newstate).

add_list(Add, Remainder, NewState) :-  append(Add, Remainder, NewState).

remove(X, [X|T], T) :- !.
remove(X, [H|T], [H|R]):- remove(X, T, R).

writeNL(X) :- write(X), nl.
writeNLNL(X) :- write(X), nl, nl.