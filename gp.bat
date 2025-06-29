@echo off
set /p MSG=Digite a mensagem do commit: 
git add .
git commit -m "%MSG%"
git push origin main
pause