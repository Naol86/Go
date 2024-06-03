import speedtest
import schedule
import time

def run_speedtest():
    st = speedtest.Speedtest()
    st.get_best_server()
    download_speed = st.download() / 1e+6  # Convert to Mbps
    upload_speed = st.upload() / 1e+6  # Convert to Mbps
    ping = st.results.ping

    print(f"Download Speed: {download_speed:.2f} Mbps")
    print(f"Upload Speed: {upload_speed:.2f} Mbps")
    print(f"Ping: {ping:.2f} ms")
    print("-" * 30)

def schedule_speedtest():
    # Run the speed test initially
    run_speedtest()

    # Schedule the speed test every minute
    schedule.every(1).minute.do(run_speedtest)

    # Run the scheduler indefinitely
    while True:
        schedule.run_pending()
        time.sleep(1)

if __name__ == "__main__":
    schedule_speedtest()
