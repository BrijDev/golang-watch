# Watch Simulator

A simple Go program that simulates a watch with mode-dependent buttons.

## Table of Contents

- [Description](#description)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Testing](#testing)

## Description

This Go program simulates a watch with two modes: "Letters" and "Numbers." It provides three buttons labeled "1," "2," and "3," and a "4" button to switch between modes. Depending on the current mode, the buttons perform different actions.

- In "Letters" mode, pressing buttons "1," "2," or "3" will print the letters "A," "B," or "C" to the console.
- In "Numbers" mode, pressing buttons "1," "2," or "3" will print the numbers 1, 2, or 3 to the console.
- Pressing button "4" toggles between "Letters" and "Numbers" mode.

This program is provided as an educational example of how to use Go structs, interfaces, and simple user input handling.

## Features

- Mode-dependent button actions.
- Mode switching between "Letters" and "Numbers."

## Prerequisites

Before running the program, make sure you have Go installed on your system.

## Getting Started

1. Clone this repository to your local machine:

   ```bash
   git clone git@github.com:BrijDev/golang-watch.git

## Usage

To run the program, execute the following command in your terminal:

```bash
go run main.go

## Testing

This program includes unit tests to verify its functionality. To run the tests, use the following command:

```bash
go test

