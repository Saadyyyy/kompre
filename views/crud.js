document.addEventListener('DOMContentLoaded', function () {
    const baseURL = 'http://localhost:8080/crud';
    
    // Fetch and display data when page loads
    fetchData();

    // Handle create form submission
    document.getElementById('createForm').addEventListener('submit', function (e) {
        e.preventDefault();

        // Construct the object to match the request body
        const kinerjaCrud = {
            id_kary: document.getElementById('idKary').value,
            nama: document.getElementById('nama').value,
            kehadiran: parseInt(document.getElementById('kehadiran').value, 10),   // Convert to integer
            jumlah_kinerjaCrud: parseInt(document.getElementById('jumlah_kinerjaCrud').value, 10),  // Convert to integer
            inisiatif: parseInt(document.getElementById('inisiatif').value, 10),  // Convert to integer
            team_work: parseInt(document.getElementById('team_work').value, 10)   // Convert to integer
        };
    
        // Call the function to create the record
        createKinerjaCrud(kinerjaCrud);
    });

    function createKinerjaCrud(kinerjaCrud) {
        fetch(`${baseURL}/create`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(kinerjaCrud),  // Convert object to JSON string
        })
        .then(response => response.json())
        .then(data => {
            console.log('Success:', data);
            // Handle success response (e.g., clear the form, show a message, etc.)
        })
        .catch((error) => {
            console.error('Error:', error);
            // Handle error (e.g., show an error message)
        });
    }

    // Function to fetch data
    function fetchData() {
        fetch(`${baseURL}/get`)
        .then(response => response.json())
        .then(data => {
            const tableBody = document.getElementById('mahasiswaTable');
            tableBody.innerHTML = '';
            data.forEach(mahasiswa => {
                const row = `<tr>
                    <td>${mahasiswa.id_kary}</td>
                    <td>${mahasiswa.nama}</td>
                    <td>${mahasiswa.kehadiran}</td>
                    <td>${mahasiswa.jumlah_kinerjaCrud}</td>
                    <td>${mahasiswa.inisiatif}</td>
                    <td>${mahasiswa.team_work}</td>
                    <td>${mahasiswa.penilaian}</td>

                    <td>
                        <button class="btn btn-warning" onclick="updateMahasiswa(${mahasiswa.ID})">Update</button>
                        <button class="btn btn-danger" onclick="deleteMahasiswa(${mahasiswa.ID})">Delete</button>
                    </td>
                </tr>`;
                tableBody.insertAdjacentHTML('beforeend', row);
            });
        })
        .catch(error => console.error('Error:', error));
    }

    window.updateMahasiswa = function(id) {
        const mahasiswa = {
            id_kary: prompt("Masukkan ID karyawan:"),
            nama: prompt("Masukkan nama mahasiswa:"),
            kehadiran: parseInt(prompt("Masukkan jumlah kehadiran:"), 10),   // Input kehadiran dari prompt
            hasilKerja: parseInt(prompt("Masukkan jumlah hasil kerja:"), 10),  // Input hasil kerja
            jumlah_kinerjaCrud: parseInt(prompt("Masukkan jumlah kinerja CRUD:"), 10),  // Input kinerja CRUD dari prompt
            inisiatif: parseInt(prompt("Masukkan nilai inisiatif:"), 10),  // Input inisiatif dari prompt
            team_work: parseInt(prompt("Masukkan nilai teamwork:"), 10)   // Input teamwork dari prompt
        };
    
        fetch(`${baseURL}/update/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(mahasiswa)
        })
        .then(response => response.json())
        .then(data => {
            alert('Mahasiswa berhasil diupdate!');
            fetchData();  // Refresh table
        })
        .catch(error => console.error('Error:', error));
    };
    
    
    // Function to delete mahasiswa
    window.deleteMahasiswa = function(id) {
        if (confirm('Apakah Anda yakin ingin menghapus mahasiswa ini?')) {
            fetch(`${baseURL}/delete/${id}`, {
                method: 'DELETE'
            })
            .then(response => response.json())
            .then(data => {
                alert('Mahasiswa berhasil dihapus!');
                fetchData();  // Refresh table
            })
            .catch(error => console.error('Error:', error));
        }
    };

    window.printTable = function() {
        var printContents = document.getElementById('mahasiswaTable').outerHTML;
        var originalContents = document.body.innerHTML;

        document.body.innerHTML = `<table class="table table-bordered">${printContents}</table>`;
        window.print();
        document.body.innerHTML = originalContents;
    }
});
