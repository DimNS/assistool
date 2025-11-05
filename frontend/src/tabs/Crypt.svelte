<script>
    import Fa from 'svelte-fa';
    import {faPlay} from '@fortawesome/free-solid-svg-icons';
    import {Encrypt, Decrypt} from '../../wailsjs/go/main/CryptTab';

    let secretKey = '';
    let withGCM = true;

    let decryptedString = '';
    let encryptedString = '';

    function onEncrypt() {
        Encrypt(secretKey, decryptedString, withGCM).then((/** @type {string} */ result) => (encryptedString = result));
    }

    function onDecrypt() {
        Decrypt(secretKey, encryptedString, withGCM).then((/** @type {string} */ result) => (decryptedString = result));
    }
</script>

<div class="mb-3">
    <label class="form-label" for="secretKey"><b>Secret key</b></label>
    <input type="text" class="form-control" id="secretKey" bind:value={secretKey} placeholder="Enter secret key" />
</div>
<div class="form-check form-switch mb-3">
    <input class="form-check-input" type="checkbox" role="switch" id="withGCM" bind:checked={withGCM} />
    <label class="form-check-label" for="withGCM">With GCM</label>
</div>

<div class="row">
    <div class="col">
        <div class="card">
            <div class="card-header">Encrypt</div>
            <div class="card-body">
                <div class="mb-3">
                    <label class="form-label" for="decryptedString"><b>Decrypted string</b></label>
                    <input
                        type="text"
                        class="form-control"
                        id="decryptedString"
                        bind:value={decryptedString}
                        placeholder="Enter decrypted string"
                    />
                </div>
                <div class="mb-3">
                    <button type="button" class="btn btn-primary" on:click={onEncrypt}>
                        <Fa icon={faPlay} />
                        Encrypt
                    </button>
                </div>
            </div>
        </div>
    </div>
    <div class="col">
        <div class="card">
            <div class="card-header">Decrypt</div>
            <div class="card-body">
                <div class="mb-3">
                    <label class="form-label" for="encryptedString"><b>Encrypted string</b></label>
                    <input
                        type="text"
                        class="form-control"
                        id="encryptedString"
                        bind:value={encryptedString}
                        placeholder="Enter encrypted string"
                    />
                </div>
                <div class="mb-3">
                    <button type="button" class="btn btn-primary" on:click={onDecrypt}>
                        <Fa icon={faPlay} />
                        Decrypt
                    </button>
                </div>
            </div>
        </div>
    </div>
</div>
