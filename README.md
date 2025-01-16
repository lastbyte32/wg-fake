# wg-fake

Fake handshake for WireGuard.
Allows to bypass DPI blocking of the WireGuard protocol by sending a "magic" packet.


### Binaries

Pre-built binaries are available [here](https://github.com/lastbyte32/wg-fake/releases/latest).

### Build from source

Run in the source directory:

```bash
$ git clone https://github.com/lastbyte32/wg-fake.git
$ make
```

Binary will be available in the `build` directory.

### Usage
In the WireGuard client configuration file, set a fixed ***ListenPort***
```
ListenPort = <LOCAL WG PORT>
```
and then right before connection start run command:

```bash
$ ./wg-fake -s <WG SERVER ADDRESS:PORT> -p <LOCAL WG PORT>
```
After successful result, you can use WireGuard client to connect to the server.
> If you use wg-quick it may be convenient to add wg-fake invocation as a PreUp command in your client config.
