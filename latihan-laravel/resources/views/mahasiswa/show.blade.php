@extends('layouts.app')

@section('konten')
<div class="container">
    <h1>Detail Mahasiswa</h1>

    <div class="card mb-4">
        <div class="card-body">
            <h4>{{ $mahasiswa->nama }}</h4>

            <p><strong>NIM:</strong> {{ $mahasiswa->nim }}</p>
            <p><strong>Email:</strong> {{ $mahasiswa->email }}</p>
            <p><strong>Program Studi:</strong> {{ $mahasiswa->programStudi->nama }}</p>
            <p><strong>Angkatan:</strong> {{ $mahasiswa->angkatan }}</p>
            <p><strong>IPK:</strong> {{ $mahasiswa->ipk }}</p>
        </div>
    </div>

    <h2>Mata Kuliah yang Diambil</h2>

    <table class="table table-bordered">
        <thead>
            <tr>
                <th>No</th>
                <th>Kode</th>
                <th>Nama Mata Kuliah</th>
                <th>SKS</th>
                <th>Semester</th>
                <th>Nilai</th>
            </tr>
        </thead>

        <tbody>
            @forelse ($mahasiswa->matakuliah as $index => $mk)
                <tr>
                    <td>{{ $index + 1 }}</td>
                    <td>{{ $mk->kode }}</td>
                    <td>{{ $mk->nama }}</td>
                    <td>{{ $mk->sks }}</td>
                    <td>{{ $mk->semester }}</td>
                    <td>{{ $mk->pivot->nilai }}</td>
                </tr>
            @empty
                <tr>
                    <td colspan="6" class="text-center">
                        Belum mengambil mata kuliah.
                    </td>
                </tr>
            @endforelse
        </tbody>
    </table>
</div>
@endsection
