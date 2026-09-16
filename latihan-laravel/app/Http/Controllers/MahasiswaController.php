<?php

namespace App\Http\Controllers;

use App\Models\Mahasiswa;

class MahasiswaController extends Controller
{
    public function show($id)
    {
        $mahasiswa = Mahasiswa::with(['programStudi', 'matakuliah'])
            ->findOrFail($id);

        return view('mahasiswa.show', compact('mahasiswa'));
    }
}
