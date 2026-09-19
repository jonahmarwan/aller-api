defmodule TCPServer do
  def start(port) do
    {:ok, socket} = :gen_tcp.listen(port, [:binary, packet: :line, active: false, reuseaddr: true])
    IO.puts("Server listening on port #{port}")
    accept_loop(socket)
  end

  defp accept_loop(socket) do
    {:ok, client_socket} = :gen_tcp.accept(socket)
    spawn(fn -> handle_client(client_socket) end)
    accept_loop(socket)
  end

  defp handle_client(socket) do
    case :gen_tcp.recv(socket, 0) do
      {:ok, data} ->
        IO.puts("Received: #{data}")
        :gen_tcp.send(socket, "Echo: #{data}")
        handle_client(socket)

      {:error, :closed} ->
        IO.puts("Client disconnected")
        :ok
    end
  end
end

TCPServer.start(4000)
