@echo off
REM Ensure Python is installed and added to the system PATH
python convert.py

REM Display the generated SQL
type output.sql

exit /b
