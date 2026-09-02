use ed25519_dalek::SigningKey;
use rand::rngs::OsRng;

#[unsafe(no_mangle)]
pub extern "C" fn gen_key(priv: *mut [u8; 32], pub: *mut[u8, 32]) {
    if priv.is_null() || pub.is_null() {
        return;}
    let mut csprng = OsRng;
    let s_key = SigningKey::generate(&mut csprng);

    let verif_key = s_key.verifying_key();

    unsafe {
        *priv = s_key.to_bytes()
        *pub = verif_key.to_bytes()
        }
}
