# Service de vote AI30 - IA04
### Amaury Michel, Pierre Gibertini
***

### Utilisation
#### 1. avec `go install`
`go install -v gitlab.utc.fr/pgiberti/ia04@latest` pour installer

`$GOPATH/bin/ia04` pour lancer le serveur sur le port local 8080


#### 2. avec `git`
`git clone https://gitlab.utc.fr/pgiberti/ia04` pour cloner le dépôt

`go run launch_ballotagent.go` pour lancer le serveur sur le port local 8080


### Travail effectué
- Implémentation complète de l'API demandée, avec toutes les vérifications des requêtes nécessaires
- Implémentation des méthodes de vote suivantes : `majority`, `borda` et `approval`


### Choix d'implémentation
- Une châine de caractère vide (`""`) dans le champ `deadline` lors d'une requête `/new_ballot` est accepté, afin de réaliser plus aisément des tests. Dans ce cas-ci, on peut accéder au résultat à n'importe quel moment, et continuer de voter après avoir accédé aux résultats.
- Si la date donnée pour le champ `deadline` lors d'une requête `/new_ballot` est antérieure à `date.Now()`, le ballot n'est pas créé.
- Tous les votants inscrits ne sont pas obligés de voter. Le résultat final dépend d'uniquement des agents ayant voté.
- En cas d'égalité, un *tie-break* par défaut a été mis en place : c'est le candidat au numéro le plus petit qui est élu.
- Le `ranking` donnée par une requête `result` affiche les candidats à égalité dans un ordre aléatoire, mais forcément après les candidats avec plus de voix. *En y repensant, il aurait été plus judicieux de tenir compte de la méthode de tie-break*.
- Le script permttant de lancer les voteurs est très basique et peu fonctionnel car nous avons réalisé le plupart de nos tests avec l'extension `RESTED` de `Mozilla Firefox`
