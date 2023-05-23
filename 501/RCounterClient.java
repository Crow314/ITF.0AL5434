//
// RCounterClient.java
//

import java.rmi.registry.LocateRegistry;
import java.rmi.registry.Registry;

class RCounterClient
{
    public static void main(String argv[]) throws Exception
    {
        if( argv.length != 2 )
        {
	    System.err.println("Usage% java RCounterClient hostname rmiregistry-portno");
	    System.exit( 1 );
        }
        String hostname = argv[0];
        String rmiregport_s = argv[1];
	int rmiregport = Integer.parseInt( rmiregport_s );
	Registry registry = LocateRegistry.getRegistry( hostname, rmiregport );

	String name = "/Counter/c1";
        IRCounter c1 = (IRCounter) registry.lookup( name );
	for( int i=0 ; i<3 ; i++ )
	{
	    System.out.println("c1.addn(3): "+c1.addn(3));
	    System.out.println("c1.value=="+c1.getValue());
	}
    }
};
