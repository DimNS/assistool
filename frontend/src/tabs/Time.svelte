<script lang="ts">
    import Fa from 'svelte-fa';
    import {faCopy, faClock} from '@fortawesome/free-solid-svg-icons';
    import {onMount, onDestroy} from 'svelte';
    import {toastSuccess} from '../services/Toast';

    let unixTimestamp = '';
    let isoString = '';
    let utcTime = '';
    let intervalId: number | undefined = undefined;

    function updateTime() {
        const now = new Date();
        unixTimestamp = Math.floor(now.getTime() / 1000).toString();
        isoString = now.toISOString().replace('Z', '+00:00');
        utcTime = now.toUTCString();
    }

    onMount(() => {
        updateTime();
        intervalId = setInterval(updateTime, 1000);
    });

    onDestroy(() => {
        clearInterval(intervalId);
    });

    function copyUnix() {
        navigator.clipboard.writeText(unixTimestamp);
        toastSuccess('Unix timestamp copied to clipboard!', true);
    }

    function copyISO() {
        navigator.clipboard.writeText(isoString);
        toastSuccess('ISO string copied to clipboard!', true);
    }
</script>

<div class="time-container">
    <div class="card mb-3">
        <div class="card-header">
            <Fa icon={faClock} />
            UTC Time
        </div>
        <div class="card-body text-center">
            <h2 class="utc-time">{utcTime}</h2>
        </div>
    </div>

    <div class="row">
        <div class="col-md-6 mb-3">
            <div class="card">
                <div class="card-header d-flex justify-content-between align-items-center">
                    Unix Timestamp
                    <button type="button" class="btn btn-outline-secondary btn-sm" on:click={copyUnix}>
                        <Fa icon={faCopy} />
                    </button>
                </div>
                <div class="card-body">
                    <code class="timestamp-value">{unixTimestamp}</code>
                </div>
            </div>
        </div>
        <div class="col-md-6 mb-3">
            <div class="card">
                <div class="card-header d-flex justify-content-between align-items-center">
                    ISO 8601
                    <button type="button" class="btn btn-outline-secondary btn-sm" on:click={copyISO}>
                        <Fa icon={faCopy} />
                    </button>
                </div>
                <div class="card-body">
                    <code class="timestamp-value">{isoString}</code>
                </div>
            </div>
        </div>
    </div>
</div>

<style>
    .utc-time {
        font-family: monospace;
        margin: 0;
    }
    .timestamp-value {
        font-size: 1.2em;
        word-break: break-all;
    }
</style>
