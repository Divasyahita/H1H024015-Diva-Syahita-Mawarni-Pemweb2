<?php

namespace Database\Seeders;

use App\Models\Matakuliah;
use Illuminate\Database\Seeder;

class MatakuliahSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        $daftarMatakuliah = [
            [
                'kode' => 'TK201',
                'nama' => 'Pemrograman Web II',
                'sks' => 3,
                'semester' => 4,
            ],
            [
                'kode' => 'TK202',
                'nama' => 'Basis Data',
                'sks' => 3,
                'semester' => 3,
            ],
            [
                'kode' => 'TK203',
                'nama' => 'Jaringan Komputer',
                'sks' => 3,
                'semester' => 3,
            ],
            [
                'kode' => 'TK204',
                'nama' => 'Sistem Operasi',
                'sks' => 3,
                'semester' => 3,
            ],
            [
                'kode' => 'TK205',
                'nama' => 'Internet of Things',
                'sks' => 2,
                'semester' => 4,
            ],
            [
                'kode' => 'TK206',
                'nama' => 'Sistem Kendali',
                'sks' => 2,
                'semester' => 4,
            ],
        ];

        foreach ($daftarMatakuliah as $matakuliah) {
            Matakuliah::create($matakuliah);
        }
    }
}