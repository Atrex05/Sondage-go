# Cahier des charges applications Sondage Web

## 1. Introduction
-----------------

### 1.1 Objectif du projet

Le but de ce projet est de réaliser une application de sondage en ligne qui permettra de créer des sondages et d'y répondre.

### 1.2 Public visé

Cette application s'adresse à toute personne souhaitant réaliser un sondage en ligne.

## 2. Description générale
-----------------

### 2.1 Fonctionnalités principales

- Création de sondage
- Répondre à un sondage
- Afficher les résultats d'un sondage
- Gestion des utilisateurs

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

### 2.5 Sécurité

- Gestion des mots de passe
- Gestion des injections SQL
- Gestion des attaques XSS
- Gestion des attaques CSRF
- Gestion des attaques par force brute

## 3. Fonctionnalités détaillées
-------------------------------

### 3.1 Création de sondage

L'utilisateur peut créer un sondage en ligne. Il doit renseigner un titre, une description, une date de début et une date de fin.

### 3.2 Répondre à un sondage

L'utilisateur peut répondre à un sondage en ligne. Il doit être connecté.

### 3.3 Afficher les résultats d'un sondage

L'utilisateur peut afficher les résultats d'un sondage en ligne.

### 3.4 Gestion des utilisateurs

L'utilisateur peut se connecter à son compte. Il doit renseigner son adresse mail et son mot de passe.

L'utilisateur peut créer un compte. Il doit renseigner son adresse mail et son mot de passe.

L'utilisateur peut se déconnecter en cliquant sur un bouton de déconnexion.

## 5. Interface utilisateur
--------------------------

### 5.1 Page d'accueil

L'utilisateur arrive sur cette page au lancement. Il peut se connecter, créer un compte ou afficher les sondages.

### 5.2 Page de connexion

L'utilisateur arrive sur cette page lorsqu'il clique sur le bouton de connexion. Il doit renseigner son adresse mail et son mot de passe.

### 5.3 Page de création de compte

L'utilisateur arrive sur cette page lorsqu'il clique sur le bouton de création de compte. Il doit renseigner son adresse mail et son mot de passe.

### 5.4 Page de création de sondage

L'utilisateur arrive sur cette page lorsqu'il clique sur le bouton de création de sondage. Il doit renseigner un titre, une description, une date de début et une date de fin.

### 5.5 Page de réponse à un sondage

L'utilisateur arrive sur cette page lorsqu'il clique sur le bouton de réponse à un sondage. Il doit renseigner ses réponses.

### 5.6 Page d'affichage des résultats d'un sondage

L'utilisateur arrive sur cette page lorsqu'il clique sur le bouton d'affichage des résultats d'un sondage. Il peut voir les résultats du sondage.

### 5.7 Page de déconnexion

L'utilisateur arrive sur cette page lorsqu'il clique sur le bouton de déconnexion. Il est déconnecté et renvoyé à la page d'accueil.

## 6. Plan de développement
--------------------------

### 6.1 Planning prévisionnel

- Semaine 1 : Création de la base de données permettant de stocker des sondages et les réponses associées.

- Semaine 2 : Création de l interface utilisateur.

- Semaine 3 : Création des fonctionnalités poussé de l application (connexion, création de compte, création de sondage, réponse à un sondage).

- Semaine 4 : Tests et débogage.

6.2 Budget: N/A
                          -Semaine 2 : Création de l interface utilisateur 
                          -Semaine 3 : Création des fonctionnalités poussé de l application (connexion, création de compte, création de sondage, réponse à un sondage)
                          -Semaine 4 : Tests et débogage

6.2 Budget: N/A

