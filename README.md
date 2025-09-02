# 💬 Net-Cat

> **Messagerie entre terminaux** - Une solution de chat en temps réel via TCP

## 🚀 Installation

Clonez le repository et utilisez le Makefile pour commencer :

```bash
git clone <repository-url>
cd net-cat
```

## 📋 Commandes disponibles

| Commande | Description |
|----------|-------------|
| `make run` | Compile le programme |
| `make clean` | Supprime les fichiers compilés |
| `make removelog` | Réinitialise le fichier logs.txt |
| `make ip` | Montre votre IP pour vous connecter au chat à distance |

## 🖥️ Utilisation

### Côté Serveur
```bash
make run
./TCPChat <port>
```
*Si aucun port n'est spécifié, le port **8989** sera utilisé par défaut.*

### Côté Client

#### 🏠 Connexion locale
```bash
nc localhost <port>
# Exemple : nc localhost 8989
```

#### 🌐 Connexion multi-ordinateurs
```bash
nc <adresse_IP> <port>
```

> ⚠️ **Attention** : Pour tester entre plusieurs ordinateurs, assurez-vous d'être sur la même connexion internet.

## ✨ Fonctionnalités

### 👤 Gestion des utilisateurs
- ✅ Choisir un nom d'utilisateur
- ✅ Changer son nom d'utilisateur
- ✅ Visualisation de tous les utilisateurs connectés

### 📺 Gestion des canaux
- ✅ Choisir un canal de discussion
- ✅ Créer un nouveau canal
- ✅ Changement de canal
- ✅ Visualisation de tous les canaux disponibles
- ✅ Historique des messages lors de la connexion à un canal

### 📊 Monitoring & Logs
- ✅ Logs automatiques (fichiers & serveur)
- ✅ Visualisation de toutes les commandes disponibles
- ✅ Notifications diverses :
  - Connexion/Déconnexion d'utilisateurs
  - Changement de nom d'utilisateur
  - Création de canaux
  - Gestion des erreurs client

### 🔧 Autres
- ✅ Déconnexion propre

## 📁 Structure

```
net-cat/
├── assets/
│   └── logs.txt          # Fichier de logs (persistant)
├── Makefile
└── ...
```

> 📝 **Note** : Le fichier `logs.txt` se trouve dans le répertoire `assets/` et ne se réinitialise que manuellement avec `make removelog`.

## 👥 Crédits

Développé avec ❤️ par :
- **Lucas Dunis** (ldunis)
- **Tristan Dumoulin** (tdumouli)

---