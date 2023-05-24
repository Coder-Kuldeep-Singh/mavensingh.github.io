import sys
import tkinter as tk
from tkinter import ttk
import os
import win32com.client


def get_battery_health_linux():
    base_path = '/sys/class/power_supply/BAT0/'
    try:
        with open(os.path.join(base_path, 'energy_full'), 'r') as f:
            energy_full = int(f.read())
        with open(os.path.join(base_path, 'energy_full_design'), 'r') as f:
            energy_full_design = int(f.read())
        health_percentage = (energy_full / energy_full_design) * 100
        return [("Battery Health", f'{health_percentage}%')]
    except FileNotFoundError:
        return [("Error", "No battery found")]
    except Exception as e:
        return [("Error", str(e))]



def get_battery_health_windows():
    strComputer = "."
    objWMIService = win32com.client.Dispatch("WbemScripting.SWbemLocator")
    objSWbemServices = objWMIService.ConnectServer(strComputer, "root\\cimv2")
    colItems = objSWbemServices.ExecQuery("SELECT * FROM Win32_Battery")
    for objItem in colItems:
        return [
            ("Availability", objItem.Availability),
            ("Battery Status", objItem.BatteryStatus),
            ("Estimated Charge Remaining", objItem.EstimatedChargeRemaining),
            ("Estimated Run Time", objItem.EstimatedRunTime),
            ("Expected Battery Life", objItem.ExpectedBatteryLife),
            ("Expected Life", objItem.ExpectedLife),
            ("Full Charge Capacity", objItem.FullChargeCapacity),
            ("Max Recharge Time", objItem.MaxRechargeTime),
            ("Status", objItem.Status),
            ("Time On Battery", objItem.TimeOnBattery),
            ("Time To Full Charge", objItem.TimeToFullCharge),
        ]
    # Return a message if no battery is found
    return [("Error", "No battery found")]



def get_battery_info():
    if sys.platform.startswith('linux'):
        try:
            return get_battery_health_linux()
        except Exception as e:
            return [("Error", str(e))]
    elif sys.platform.startswith('win32'):
        try:
            return get_battery_health_windows()
        except Exception as e:
            return [("Error", str(e))]
    else:
        return [("Error", "OS not supported")]



def refresh_table(tree):
    for i in tree.get_children():
        tree.delete(i)

    for name, value in get_battery_info():
        tree.insert('', 'end', values=(name, value))

    # Refresh the table every second
    tree.after(1000, refresh_table, tree)


root = tk.Tk()
tree = ttk.Treeview(root, columns=("Name", "Value"), show="headings")
tree.heading("Name", text="Name")
tree.heading("Value", text="Value")
tree.pack()
refresh_table(tree)
root.mainloop()
