use ed25519_dalek::SigningKey;
use rand::rngs::SysRng;
use rand_core::UnwrapErr;

#[unsafe(no_mangle)]
pub extern "C" fn gen_key(priv_key: *mut [u8; 32], pub_key: *mut[u8; 32]) {
    if priv_key.is_null() || pub_key.is_null() {
        return;}
    let mut csprng = UnwrapErr(SysRng);
    let s_key = SigningKey::generate(&mut csprng);

    let verif_key = s_key.verifying_key();

    unsafe {
        *priv_key = s_key.to_bytes();
        *pub_key = verif_key.to_bytes();
        }
}
