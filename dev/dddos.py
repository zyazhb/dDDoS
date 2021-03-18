#!/usr/bin/env python

from mininet.net import Mininet
from mininet.node import Controller, RemoteController, OVSController
from mininet.node import CPULimitedHost, Host, Node
from mininet.node import OVSKernelSwitch, UserSwitch
from mininet.node import IVSSwitch
from mininet.cli import CLI
from mininet.log import setLogLevel, info
from mininet.link import TCLink, Intf
from subprocess import call

def myNetwork():

    net = Mininet( topo=None,
                   build=False,
                   ipBase='10.0.0.0/8')

    info( '*** Adding controller\n' )
    c0=net.addController(name='c0',
                      controller=Controller,
                      protocol='tcp',
                      port=6633)

    info( '*** Add switches\n')
    s22 = net.addSwitch('s22', cls=OVSKernelSwitch)
    s9 = net.addSwitch('s9', cls=OVSKernelSwitch)
    s23 = net.addSwitch('s23', cls=OVSKernelSwitch)
    s21 = net.addSwitch('s21', cls=OVSKernelSwitch)

    info( '*** Add hosts\n')
    h16 = net.addHost('h16', cls=Host, ip='10.0.1.16', defaultRoute=None)
    h23 = net.addHost('h23', cls=Host, ip='10.0.0.23', defaultRoute=None)
    e1 = net.addHost('e1', cls=Host, ip='10.0.0.4', defaultRoute=None)
    h21 = net.addHost('h21', cls=Host, ip='10.0.0.21', defaultRoute=None)
    h25 = net.addHost('h25', cls=Host, ip='10.0.0.25', defaultRoute=None)
    h1 = net.addHost('h1', cls=Host, ip='10.0.0.1', defaultRoute=None)
    h18 = net.addHost('h18', cls=Host, ip='10.0.0.18', defaultRoute=None)
    e3 = net.addHost('e3', cls=Host, ip='10.0.0.19', defaultRoute=None)
    h26 = net.addHost('h26', cls=Host, ip='10.0.0.26', defaultRoute=None)
    h2 = net.addHost('h2', cls=Host, ip='10.0.0.2', defaultRoute=None)
    e4 = net.addHost('e4', cls=Host, ip='10.0.0.20', defaultRoute=None)
    h3 = net.addHost('h3', cls=Host, ip='10.0.0.3', defaultRoute=None)
    h24 = net.addHost('h24', cls=Host, ip='10.0.0.24', defaultRoute=None)
    h22 = net.addHost('h22', cls=Host, ip='10.0.0.22', defaultRoute=None)
    h17 = net.addHost('h17', cls=Host, ip='10.0.0.17', defaultRoute=None)
    e2 = net.addHost('e2', cls=Host, ip='10.0.1.15', defaultRoute=None)

    info( '*** Add links\n')
    net.addLink(h16, s21)
    net.addLink(s21, e2)
    net.addLink(h17, s21)
    net.addLink(h18, s21)
    net.addLink(s21, s22)
    net.addLink(s22, s23)
    net.addLink(h21, s22)
    net.addLink(s22, e3)
    net.addLink(h22, s23)
    net.addLink(s23, e4)
    net.addLink(s9, s21)
    net.addLink(h23, s22)
    net.addLink(h24, s22)
    net.addLink(h25, s23)
    net.addLink(h26, s23)
    net.addLink(h1, s9)
    net.addLink(h2, s9)
    net.addLink(h3, s9)
    net.addLink(s9, e1)

    info( '*** Starting network\n')
    net.build()
    info( '*** Starting controllers\n')
    for controller in net.controllers:
        controller.start()

    info( '*** Starting switches\n')
    net.get('s22').start([c0])
    net.get('s9').start([c0])
    net.get('s23').start([c0])
    net.get('s21').start([c0])

    info( '*** Post configure switches and hosts\n')
    s9.cmd('ifconfig s9 10.0.0.1')

    CLI(net)
    net.stop()

if __name__ == '__main__':
    setLogLevel( 'info' )
    myNetwork()

