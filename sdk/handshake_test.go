package sdk

import "testing"

func TestHandshakeConfig(t *testing.T) {
	if HandshakeConfig.ProtocolVersion != ProtocolVersion {
		t.Fatalf("unexpected protocol version: %d", HandshakeConfig.ProtocolVersion)
	}
	if HandshakeConfig.MagicCookieKey != MagicCookieKey {
		t.Fatalf("unexpected cookie key: %s", HandshakeConfig.MagicCookieKey)
	}
	if HandshakeConfig.MagicCookieValue != MagicCookieValue {
		t.Fatalf("unexpected cookie value: %s", HandshakeConfig.MagicCookieValue)
	}
}
