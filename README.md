# SmartHome controller

A smart home controller is a service that provides an interface for monitoring the state of home systems.
It receives information from sensors, stores it, and can provide information upon request.
The controller's functionality works with two types of sensors:
- `ContactClosure` - dry contacts (leakage sensor, switch for closing or opening).
  Events usually arrive when the sensor state changes. For example, when a door or window is opened.
- `ADC` - analog input (thermometers, level sensors). Events usually arrive when sensor readings change.
  For example, when temperature or humidity sensor readings change.

_The project was written on the educational course T-Education._