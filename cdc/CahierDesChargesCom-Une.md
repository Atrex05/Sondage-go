# Cahier des charges Com-Une

## 1. Introduction
-----------------

### 1.1 Objectif du projet
Améliorer la communication, l organisation et la gestion des ressources lors d une crise sanitaire.
Pour cela, nous allons réaliser une simulation de scénario catastrophe dans un hôpital.
L'objectif sera d'être le plus réaliste possible afin de pouvoir former le personnel hôspitalier.

### 1.2 Public visé
DMC, Directeurs, Cadres paramedicaux, régulateurs SAMU, COP, COS, Médecins
--> Personnel hospitalier

## 2. Description générale
-----------------

### 2.1 Fonctionnalités principales
#### Maître du jeu:
Le maître du jeu est la personne qui va créer la simulation. Son rôle va être la mise en place contextuelle de la simulation.
##### Possibilité de créer une simulation réaliste de catastrophe:
    - Création de patient via la fiche macsim
        * Lors de la création de la simulation, le maître du jeu pourra créer des patients avec des pathologies différentes
    - Création de personnel hospitalier
        * Lors de la création de la simulation, le maître du jeu pourra créer du personnel hospitalier avec des compétences différentes(spécialités métiers, nombres de prise en charge simultanée maximum, etc...)
    - Création de ressources
        * Lors de la création de la simulation, le maître du jeu pourra créer des ressources tel que du matériel médical, des médicaments, ou encore des utilitaires tel que des sacs poubelles, des draps, etc...
##### Possibilité de crée un hôpital:
    - Le maître du jeu pourra créer des hôpitaux avec des spécialités différentes
        *  Il sera possible de choisir le nombres de salles ainsi que leur répartition, le nombres de lits, le nombres de places en réanimation, etc...

#### Joueur:
Le joueur est la personne qui va jouer la simulation. Son rôle va être la bonne gestion de la simulation.
Chaque joueur sera assigné à un hôpital. Les différents joueur seront ammené à collaboré pour la bonne gestion des patients.
##### Gestion des ressources:
    - Le joueur sera ammené à gérer les ressources humaines (personnel hospitalier et patients)
    - Le joueur sera ammené à gérer des ressources matériels (matériels médicals, jetables type sac poubelle, etc...)
    Toute la reussite de la simulation dépendra de la gestion du joueur. Il devra gérer de manière précautionneuse chaque patient et chaque personnel hospitalier en veillant à ce que personne ne soit pas surchargé et que les patients soit bien pris en charge.
##### Patient:
    - Chaque patient va avoir une fiche macsim qui va permettre au joueur de savoir quelles sont les pathologies du patient, son état de santé, son état de conscience, etc... Il y aura sur la fiche macsim, une couleur indiquant la gravité du patient. Plus l'état est grave, plus il va nécesssité une intervention rapide.
##### Personnels hôspitalier:
    - Chaque personnel hospitalier aura un domaine de prédilection (cardiologie, pneumologie, etc...)
    - Un personnel hospitalier peut s'occuper d'un patient ne faisant pas partie de son domaine de prédilection mais il sera moins efficace.
    - Chaque personnel hospitalier aura un nombre de prise en charge simultanée maximum. Si ce nombre est dépassé, le personnel hospitalier sera grandement moins efficace.
    - Chaque personnel hospitalier ne pourra pas être actif plus de 6 heures consécutives. Il devra se reposer au moins X heures avant de pouvoir reprendre son activité.
##### Hôpital:
    - Chaque hôpital aura une capacité limité de prise en charge de patient. C'est au joueur d'accepter ou non la prise en charges de ces patients. Il sera possible d'envoyer certains patients dans d'autres hôpitaux cependant cette fonction sera limité à X par heures.
    - Les informations conçernant l'hôpital seront affiché sur l'interface du joueur. Il pourra voir le nombres de lits, le nombres de places en réanimation, le nombres de salles disponibles, etc...

### 2.2 Vision globale
    - Le jeu serra un "Sims like". Il se jouera avec une vue de dessus. 
    - Le joueur pourra se déplacer dans l'hôpital et interagir avec les patients et le personnel hospitalier.
    - Le temps sera accélerer 1h de jeu = X minutes dans la simulation. 
    - Le joueur aura un bouton "panic", une pause en cas d'extrême nécessité car il ne peut plus gérer le jeu, déborder etc. 
        (ne pas        l'appuyer en temps normal) aucun patient ne peut plus être gérer. Permets de faire un point pour ne pas perdre le jeu.
    - Le jeu sera multi-joueur. Chaque joueur sera assigné à un hôpital. Il pourra communiquer avec les autres joueurs via X pour une  organisation commune.

### 2.3 Exigences techniques

- Golang : Backend
- PostgreSQL : Base de données
- HTML/CSS : Frontend
- Javascript : Frontend

### 2.4 Frameworks et librairies

- Gorilla : Golang
- net/http : Golang
- database/sql : Golang
- html/template : Golang
- joncalhoun/form : Golang
- Tailwind CSS : HTML/CSS (optionnel)
- D'autres librairies seront ajoutées au fur et à mesure du développement

### 2.5 Sécurité

- Gestion des mots de passe
- Gestion des injections SQL
- Gestion des attaques XSS
- Gestion des attaques CSRF
- Gestion des attaques par force brute

## 3. Fonctionnalités détaillées
-------------------------------

### 3.1 Gestion des utilisateurs

L'utilisateur peut se connecter à son compte. Il doit renseigner son adresse mail et son mot de passe.

L'utilisateur peut créer un compte. Il doit renseigner son adresse mail et son mot de passe.

L'utilisateur peut se déconnecter en cliquant sur un bouton de déconnexion.

Certains utilisateurs auront des droits supplémentaires. Ils pourront créer des simulations.

## 4. Interface utilisateur
--------------------------

A définir

## 5 Planning prévisionnel

- Pour le 29/09 : Création d'une première version permettant l'ajout de patient et de personnel hospitalier stocké dans une base de données.

Le reste sera à définir au fur et à mesure du développement.