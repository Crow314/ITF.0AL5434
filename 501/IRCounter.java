//
// IRCounter.java
//

public interface IRCounter extends java.rmi.Remote
{
    public void	up() throws java.rmi.RemoteException;
    public int	getValue() throws java.rmi.RemoteException;
    public void	reset(int newVal) throws java.rmi.RemoteException;
}
