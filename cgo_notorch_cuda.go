//go:build cuda

// CUDA-build linkage + GPU enable for the notorch trainer. Links
// libnotorch_gpu.a (notorch.c compiled -DUSE_CUDA, plus notorch_cuda.o)
// and cuBLAS/cudart, so the notorch training tape dispatches matvecs to
// the device. ntGPUEnable() is the automatic GPU/CPU rule (06_PLAN §8);
// the !cuda counterpart is the no-op stub in gpu_notorch_stub.go.

package main

/*
#cgo linux CFLAGS: -DUSE_CUDA -I/usr/local/cuda/include
#cgo linux LDFLAGS: -L/usr/local/lib -lnotorch_gpu -L/usr/local/cuda/lib64 -lcudart -lcublas -lstdc++ -L/usr/lib/x86_64-linux-gnu/openblas-pthread/ -lopenblas -lm
#include <notorch.h>

// notorch_cuda.h pulls in the CUDA runtime headers, which the cgo C
// parser stumbles on. Forward-declare the four GPU entry points the
// trainer needs — the symbols resolve from libnotorch_gpu.a at link time.
int       gpu_init(void);
void      gpu_shutdown(void);
long long nt_gpu_dispatch_count(void);
void      nt_gpu_dispatch_reset(void);
*/
import "C"

// ntGPUEnable runs gpu_init(); on success it switches the notorch tape to
// GPU dispatch (nt_set_gpu_mode) and clears the dispatch counter. On
// failure the trainer stays on CPU/BLAS. Automatic, no flag (06_PLAN §8).
func ntGPUEnable() (bool, string) {
	if C.gpu_init() != 0 {
		return false, "gpu_init() failed — notorch trainer on CPU/BLAS"
	}
	C.nt_set_gpu_mode(0) // device inited; start on CPU — ntSetGPUForStage gates by stage
	C.nt_gpu_dispatch_reset()
	return true, "notorch GPU device initialized — stage-gated (CPU until teen, cuBLAS at teen/adult)"
}

// ntSetGPUForStage routes the notorch training tape to GPU only when the model
// is big enough to amortize cuBLAS kernel-launch overhead (teen+, NEmbd>=224).
// Below that, the tiny matmuls run faster on CPU/BLAS (measured 8 steps/s on
// GPU at child vs ~90 on CPU). Called at startup and after each ontogenesis
// growth. (2026-06-03 GPU mechanism rework.)
func ntSetGPUForStage(stage int) {
	if stage >= 4 { // GrowthStages index: 4=teen (224d), 5=adult (320d)
		C.nt_set_gpu_mode(1)
	} else {
		C.nt_set_gpu_mode(0)
	}
}

// ntGPUDispatchCount returns the cuBLAS dispatch count — criterion-4 proof
// that the training tape's matvecs reached the device.
func ntGPUDispatchCount() int64 { return int64(C.nt_gpu_dispatch_count()) }
