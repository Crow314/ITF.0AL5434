//
// RCounter.java
//

import java.rmi.*;
import java.rmi.server.*;

public class RCounter extends java.rmi.server.UnicastRemoteObject
    implements IRCounter
{
    int val;
    public RCounter(int initVal) throws RemoteException
    {
	super();
	val = initVal ;
    }
    public void up() throws java.rmi.RemoteException
    {
	val ++ ;
    }
    public int getValue() throws java.rmi.RemoteException
    {
	return( val );
    }
    public void reset(int newVal) throws java.rmi.RemoteException
    {
	val = newVal ;
    }
};

