//
// RCounterServer.java
//

import java.rmi.registry.Registry;
import java.rmi.registry.LocateRegistry;
import java.rmi.RemoteException;
import java.rmi.server.UnicastRemoteObject;

public class RCounterServer
{
    public static void main(String argv[]) throws Exception
    {
        if( argv.length != 1 )
        {
	    System.err.println("Usage% java RCounterServer rmiregistry-portno");
	    System.exit( 1 );
        }
        String rmiregport_s = argv[0];
	int rmiregport = Integer.parseInt( rmiregport_s );
	Registry registry = LocateRegistry.getRegistry( rmiregport );

	IRCounter c1 = new RCounter( 10 );
	String name = "/Counter/c1" ;
	registry.rebind( name, c1 );
    }
};
