SynthDef(\Beats, {
    var lfo = SinOsc.kr(0.2, add: 440);
    Out.ar(0, SinOsc.ar(lfo), SinOsc.ar(lfo));
}).writeDefFile(File.getcwd);

SynthDef(\foo, {
    Out.ar(0, SinOsc.ar() * Blip.ar());
}).writeDefFile(File.getcwd);

SynthDef(\bar, {
    Out.ar(0, SinOsc.ar(mul: Blip.ar()));
}).writeDefFile(File.getcwd);

SynthDef(\baz, {
    Out.ar(0, Blip.ar(mul: SinOsc.ar()));
}).writeDefFile(File.getcwd);

SynthDef(\sub, {
    Out.ar(0, SinOsc.ar() - Blip.ar());
}).writeDefFile(File.getcwd);

SynthDef(\Envgen1, {
    Out.ar(0, PinkNoise.ar() * EnvGen.kr(Env.perc, doneAction: 2));
}).writeDefFile(File.getcwd);

SynthDef(\defWith2Params, {
    arg freq=440, gain=0.5;
    var env = EnvGen.kr(Env.perc, doneAction: 2, levelScale: gain);
    var sine = SinOsc.ar(freq);
    Out.ar(0, sine * env);
}).writeDefFile(File.getcwd);

SynthDef(\SameSame, {
    var s = SinOsc.ar(220);
    Out.ar(0, [s, s]);
}).writeDefFile(File.getcwd);

SynthDef(\SawTone1, {
    arg freq=440, cutoff=1200, q=0.5;
    Out.ar(0, RLPF.ar(Saw.ar(freq), cutoff, q));
}).writeDefFile(File.getcwd);

SynthDef(\SineTone, {
    Out.ar(0, SinOsc.ar(440));
}).writeDefFile(File.getcwd);

SynthDef(\SineTone2, {
    Out.ar(0, SinOsc.ar(440, SinOsc.ar(0.1), 0.5));
}).writeDefFile(File.getcwd);

SynthDef(\SineTone3, {
    Out.ar(0, SinOsc.ar(440, SinOsc.ar(0.1), add: 0.5));
}).writeDefFile(File.getcwd);

SynthDef(\SineTone4, {
    arg freq=440;
    Out.ar(0, SinOsc.ar(freq));
}).writeDefFile(File.getcwd);

SynthDef(\UseParam, {
	arg freq=200;
	Out.ar(0, SinOsc.ar(freq + 20));
}).writeDefFile(File.getcwd);

SynthDef(\SimpleMulti, {
}).writeDefFile(File.getcwd);

SynthDef(\Cascade, {
    var mod1 = SinOsc.ar([440, 441]);
    var mod2 = SinOsc.ar(mod1);
    Out.ar(0, SinOsc.ar(mod2));
}).writeDefFile(File.getcwd);

SynthDef(\AllpassExample, {
    Out.ar(0, AllpassC.ar(WhiteNoise.ar(0.1), 0.01, XLine.kr(0.0001, 0.01, 20), 0.2));
}).writeDefFile(File.getcwd);

SynthDef(\AllpassnExample, {
    var noise = WhiteNoise.ar();
    var dust = Dust.ar(1, 0.5);
    var decay = Decay.ar(dust, 0.2, noise);
    var sig = AllpassN.ar(decay, 0.2, 0.2, 3);
    Out.ar(0, sig);
}).writeDefFile(File.getcwd);

SynthDef(\IntegratorExample, {
    Out.ar(0, Integrator.ar(LFPulse.ar(1500 / 4, 0.2, 0.1), MouseX.kr(0.01, 0.999, 1)));
}).writeDefFile(File.getcwd);

SynthDef(\FSinOscExample, {
    Out.ar(0, FSinOsc.ar(FSinOsc.ar(XLine.kr(4, 401, 8), 0.0, 200, 800)) * 0.2);
}).writeDefFile(File.getcwd);

SynthDef(\BPFExample, {
    var line = XLine.kr(0.7, 300, 20);
    var saw = Saw.ar(200, 0.5);
    var sine = FSinOsc.kr(line, 0, 3600, 4000);
    Out.ar(0, BPF.ar(saw, sine, 0.3));
}).writeDefFile(File.getcwd);

SynthDef(\BRFExample, {
    var line = XLine.kr(0.7, 300, 20);
    var saw = Saw.ar(200, 0.5);
    var sine = FSinOsc.kr(line, 0, 3800, 4000);
    Out.ar(0, BRF.ar(saw, sine, 0.3));
}).writeDefFile(File.getcwd);

SynthDef(\Balance2Example, {
    var l = LFSaw.ar(44);
    var r = Pulse.ar(33);
    var pos = FSinOsc.kr(0.5);
    var level = 0.1;
    Out.ar(0, Balance2.ar(l, r, pos, level));
}).writeDefFile(File.getcwd);

SynthDef(\BlipExample, {
    var freq = XLine.kr(20000, 200, 6);
    var harms = 100;
    var mul = 0.2;
    Out.ar(0, Blip.ar(freq, harms, mul));
}).writeDefFile(File.getcwd);

SynthDef(\LFSawExample, {
    var freq = LFSaw.kr(4, 0, 200, 400);
    Out.ar(0, LFSaw.ar(freq, 0, 0.1));
}).writeDefFile(File.getcwd);

SynthDef(\LFPulseExample, {
    var freq = LFPulse.kr(3, 0, 0.3, 200, 200);
    Out.ar(0, LFPulse.ar(freq, 0, 0.2, 0.1));
}).writeDefFile(File.getcwd);

SynthDef(\ImpulseExample, {
    var freq = XLine.kr(800, 100, 5);
    var gain = 0.5;
    var phase = 0.0;
    var sig = Impulse.ar(freq, phase, gain);
    Out.ar(0, sig);
}).writeDefFile(File.getcwd);

SynthDef(\LFNoise1Example, {
    var freq = XLine.kr(1000, 10000, 10);
    Out.ar(0, LFNoise1.ar(freq, 0.25));
}).writeDefFile(File.getcwd);

SynthDef(\LFTriExample, {
    var freq = LFTri.kr(4, 0, 200, 400);
    Out.ar(0, LFTri.ar(freq, 0, 0.1));
}).writeDefFile(File.getcwd);

SynthDef(\PlayBufExample, {
    arg bufnum = 0;
    Out.ar(0, PlayBuf.ar(1, bufnum, 1.0, 1.0, 0, 0, 2));
}).writeDefFile(File.getcwd);

SynthDef(\CrackleTest, {
    var crack = Crackle.ar(Line.kr(1.0, 2.0, 3), 0.5, 0.5);
    Out.ar(0, crack);
}).writeDefFile(File.getcwd);

SynthDef(\GrainBufTest, {
    Out.ar(0, GrainBuf.ar(numChannels: 1, sndbuf: 0));
}).writeDefFile(File.getcwd);

SynthDef(\COscTest, {
    Out.ar(0, COsc.ar(0, 200, 0.7, 0.25));
}).writeDefFile(File.getcwd);

SynthDef(\ClipNoiseTest, {
    Out.ar(0, ClipNoise.ar(0.2));
}).writeDefFile(File.getcwd);

SynthDef(\CombCTest, {
    var line = XLine.kr(0.0001, 0.01, 20);
    var sig = CombC.ar(WhiteNoise.ar(0.01), 0.01, line, 0.2);
    Out.ar(0, sig);
}).writeDefFile(File.getcwd);

SynthDef(\CombLTest, {
    var line = XLine.kr(0.0001, 0.01, 20);
    var sig = CombL.ar(WhiteNoise.ar(0.01), 0.01, line, 0.2);
    Out.ar(0, sig);
}).writeDefFile(File.getcwd);

SynthDef(\CombNTest, {
    var line = XLine.kr(0.0001, 0.01, 20);
    var sig = CombN.ar(WhiteNoise.ar(0.01), 0.01, line, 0.2);
    Out.ar(0, sig);
}).writeDefFile(File.getcwd);

0.exit;
